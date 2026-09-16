package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// AppEntry is one product on the /apps board. The registry is static: the
// set of apps changes with a deploy, never at runtime.
type AppEntry struct {
	Name        string
	Slug        string
	Description string
	// URL is the app's /internal/metrics endpoint. Empty means the app is not
	// wired up (not deployed, or no env var set) and it renders as "—".
	URL string
}

// AppMetricsResult is one app's user count, or the reason there isn't one.
type AppMetricsResult struct {
	App AppEntry
	// Users is nil when the fetch failed or the endpoint is not configured.
	// A pointer rather than a zero so that "0 users" and "unknown" stay
	// distinguishable on the page.
	Users *int64
	Err   error
}

// Reachable reports whether a count was actually obtained.
func (r AppMetricsResult) Reachable() bool { return r.Users != nil }

// ErrNotConfigured is the Err on a result whose app has no metrics URL.
var ErrNotConfigured = errors.New("metrics endpoint not configured")

const (
	// appMetricsTimeout bounds one page load. Every endpoint is fetched in
	// parallel, so this is the worst case for the whole board, not per app.
	appMetricsTimeout = 5 * time.Second
	// appMetricsTTL is how long a fetched board is served before the
	// backends are asked again. /apps is public, so without this every
	// visitor would fan out five requests into the cluster.
	appMetricsTTL = 60 * time.Second
)

// AppMetricsService fetches user counts from each app's /internal/metrics
// endpoint, in parallel, with a shared secret and a short cache.
type AppMetricsService struct {
	apps   []AppEntry
	secret string
	client *http.Client
	ttl    time.Duration

	mu        sync.Mutex
	cached    []AppMetricsResult
	fetchedAt time.Time
}

// NewAppMetricsService builds the production registry from the environment:
// METRICS_SECRET plus one <APP>_METRICS_URL per app.
func NewAppMetricsService() *AppMetricsService {
	apps := []AppEntry{
		{
			Name:        "Norviq",
			Slug:        "norviq",
			Description: "Stock plan and portfolio intelligence for equity-compensated employees.",
			URL:         os.Getenv("NORVIQ_METRICS_URL"),
		},
		{
			Name:        "Loci",
			Slug:        "loci",
			Description: "AI travel companion that builds itineraries from local context.",
			URL:         os.Getenv("LOCI_METRICS_URL"),
		},
		{
			Name:        "Seshat",
			Slug:        "seshat",
			Description: "Interactive coding lessons with live in-browser execution.",
			URL:         os.Getenv("SESHAT_METRICS_URL"),
		},
		{
			Name:        "Skyvisor",
			Slug:        "skyvisor",
			Description: "Operational intelligence for aviation professionals.",
			URL:         os.Getenv("SKYVISOR_METRICS_URL"),
		},
		{
			Name:        "North",
			Slug:        "north",
			Description: "AI health coach for training, nutrition, and recovery.",
			URL:         os.Getenv("NORTH_METRICS_URL"),
		},
	}
	return NewAppMetricsServiceWith(apps, os.Getenv("METRICS_SECRET"), nil)
}

// NewAppMetricsServiceWith is the constructor tests use: an explicit registry,
// secret and client. A nil client gets the production timeout.
func NewAppMetricsServiceWith(apps []AppEntry, secret string, client *http.Client) *AppMetricsService {
	if client == nil {
		client = &http.Client{Timeout: appMetricsTimeout}
	}
	return &AppMetricsService{
		apps:   apps,
		secret: secret,
		client: client,
		ttl:    appMetricsTTL,
	}
}

// SetTTL overrides the cache lifetime. Zero disables caching.
func (s *AppMetricsService) SetTTL(ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ttl = ttl
	s.cached = nil
}

// Apps returns the registry with every count unknown. It is what a nil
// service renders, so the page has a shape even before anything is wired up.
func (s *AppMetricsService) Apps() []AppMetricsResult {
	if s == nil {
		return nil
	}
	out := make([]AppMetricsResult, len(s.apps))
	for i, app := range s.apps {
		out[i] = AppMetricsResult{App: app, Err: ErrNotConfigured}
	}
	return out
}

// FetchAll returns every app's count, fetching in parallel when the cache is
// stale. Results keep registry order. A nil receiver is a legal empty board.
func (s *AppMetricsService) FetchAll(ctx context.Context) []AppMetricsResult {
	if s == nil {
		return nil
	}

	s.mu.Lock()
	if s.cached != nil && s.ttl > 0 && time.Since(s.fetchedAt) < s.ttl {
		cached := s.cached
		s.mu.Unlock()
		return cached
	}
	s.mu.Unlock()

	ctx, cancel := context.WithTimeout(ctx, appMetricsTimeout)
	defer cancel()

	results := make([]AppMetricsResult, len(s.apps))
	var wg sync.WaitGroup
	for i, app := range s.apps {
		results[i] = AppMetricsResult{App: app}
		if app.URL == "" {
			results[i].Err = ErrNotConfigured
			continue
		}
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			count, err := s.fetchOne(ctx, url)
			if err != nil {
				results[i].Err = err
				return
			}
			results[i].Users = &count
		}(i, app.URL)
	}
	wg.Wait()

	s.mu.Lock()
	s.cached = results
	s.fetchedAt = time.Now()
	s.mu.Unlock()
	return results
}

func (s *AppMetricsService) fetchOne(ctx context.Context, url string) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("build request: %w", err)
	}
	if s.secret != "" {
		req.Header.Set("Authorization", "Bearer "+s.secret)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("fetch: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status %d", resp.StatusCode)
	}

	var payload struct {
		Users *int64 `json:"users"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return 0, fmt.Errorf("decode: %w", err)
	}
	if payload.Users == nil {
		return 0, errors.New("decode: missing users field")
	}
	return *payload.Users, nil
}

// TotalUsers sums the counts that were obtained; unreachable apps add nothing.
func TotalUsers(results []AppMetricsResult) int64 {
	var total int64
	for _, r := range results {
		if r.Users != nil {
			total += *r.Users
		}
	}
	return total
}

// ReportingCount is how many apps answered.
func ReportingCount(results []AppMetricsResult) int {
	n := 0
	for _, r := range results {
		if r.Users != nil {
			n++
		}
	}
	return n
}

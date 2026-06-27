package service

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

type BlogFlagStat struct {
	Code  string
	Count int
}

type BlogStatsSnapshot struct {
	TotalViews int
	Flags      []BlogFlagStat
}

type BlogTracker struct {
	mu       sync.Mutex
	path     string
	total    int
	unique   map[string]bool
	flags    map[string]int
	client   *http.Client
	disabled bool
}

type blogStatsFile struct {
	TotalViews int            `json:"totalViews"`
	UniqueIPs  []string       `json:"uniqueIPs"`
	Flags      map[string]int `json:"flags"`
}

func NewBlogTracker() *BlogTracker {
	tracker := &BlogTracker{
		path:   resolveBlogStatsPath(),
		unique: map[string]bool{},
		flags:  map[string]int{},
		client: &http.Client{Timeout: 2 * time.Second},
	}
	tracker.load()
	return tracker
}

func (t *BlogTracker) RecordVisit(r *http.Request) {
	if t == nil || t.disabled {
		return
	}

	ip := clientIP(r)

	t.mu.Lock()
	t.total++
	isNewIP := ip != "" && !t.unique[ip]
	if isNewIP {
		t.unique[ip] = true
	}
	t.saveLocked()
	t.mu.Unlock()

	if !isNewIP {
		return
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return
	}
	if parsedIP.IsLoopback() {
		t.addCountryCode("US")
		return
	}
	if parsedIP.IsPrivate() {
		return
	}

	go t.lookupAndAddCountryCode(ip)
}

func (t *BlogTracker) Snapshot() BlogStatsSnapshot {
	if t == nil {
		return BlogStatsSnapshot{Flags: []BlogFlagStat{}}
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	flags := make([]BlogFlagStat, 0, len(t.flags))
	for code, count := range t.flags {
		flags = append(flags, BlogFlagStat{
			Code:  code,
			Count: count,
		})
	}
	sort.Slice(flags, func(i, j int) bool {
		return flags[i].Count > flags[j].Count
	})

	return BlogStatsSnapshot{
		TotalViews: t.total,
		Flags:      flags,
	}
}

func (t *BlogTracker) load() {
	data, err := os.ReadFile(t.path)
	if err != nil {
		return
	}

	var file blogStatsFile
	if err := json.Unmarshal(data, &file); err != nil {
		t.disabled = true
		return
	}

	t.total = file.TotalViews
	for _, ip := range file.UniqueIPs {
		t.unique[ip] = true
	}
	if file.Flags != nil {
		t.flags = file.Flags
	}
}

func (t *BlogTracker) lookupAndAddCountryCode(ip string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://ip-api.com/json/"+ip+"?fields=status,countryCode", nil)
	if err != nil {
		return
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var decoded struct {
		Status      string `json:"status"`
		CountryCode string `json:"countryCode"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return
	}
	if decoded.Status != "success" || decoded.CountryCode == "" {
		return
	}

	t.addCountryCode(decoded.CountryCode)
}

func (t *BlogTracker) addCountryCode(code string) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if len(code) != 2 {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.flags[code]++
	t.saveLocked()
}

func (t *BlogTracker) saveLocked() {
	if err := os.MkdirAll(filepath.Dir(t.path), 0o755); err != nil {
		t.disabled = true
		return
	}

	uniqueIPs := make([]string, 0, len(t.unique))
	for ip := range t.unique {
		uniqueIPs = append(uniqueIPs, ip)
	}
	sort.Strings(uniqueIPs)

	data, err := json.Marshal(blogStatsFile{
		TotalViews: t.total,
		UniqueIPs:  uniqueIPs,
		Flags:      t.flags,
	})
	if err != nil {
		t.disabled = true
		return
	}
	if err := os.WriteFile(t.path, data, 0o644); err != nil {
		t.disabled = true
	}
}

func resolveBlogStatsPath() string {
	if explicitPath := os.Getenv("BLOG_STATS_PATH"); explicitPath != "" {
		return explicitPath
	}
	if dbPath := os.Getenv("DATABASE_PATH"); dbPath != "" {
		return filepath.Join(filepath.Dir(dbPath), "blog_stats.json")
	}
	return "blog_stats.json"
}

func clientIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}
	if realIP := strings.TrimSpace(r.Header.Get("X-Real-IP")); realIP != "" {
		return realIP
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	return strings.TrimSpace(r.RemoteAddr)
}

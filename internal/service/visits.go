package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"myapp/internal/db"
)

// Window sizes for everything /stats reports. The retention sweep uses the
// same window, so the table and the globe always agree.
const (
	statsWindow    = "-24 hours"
	liveWindow     = 5 * time.Minute
	ringCapacity   = 200
	maxSubscribers = 100
	snapshotTTL    = 15 * time.Second
)

// VisitPing is one page view, already stripped of anything identifying. This
// is the exact shape sent over SSE and embedded in the initial globe payload.
type VisitPing struct {
	CountryCode string    `json:"cc"`
	City        string    `json:"city"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	Path        string    `json:"path"`
	At          time.Time `json:"at"`
}

type CountryCount struct {
	Code   string
	Flag   string
	Visits int
}

type PathCount struct {
	Path   string
	Visits int
}

// StatsSnapshot is the server-rendered half of /stats and the ops strip on the
// landing page. Everything in it is derived from real traffic.
type StatsSnapshot struct {
	VisitorsNow  int
	Views24h     int
	Visitors24h  int
	Countries24h int
	TopCountries []CountryCount
	TopPaths     []PathCount
	Uptime       time.Duration
	ServingFrom  string
}

type geoResult struct {
	countryCode string
	city        string
	lat         float64
	lon         float64
}

// VisitTracker records page views for the /stats globe.
//
// It deliberately keeps no personal data. The IP is salted and hashed on the
// way in and then discarded; geo lookups are keyed by network prefix rather
// than by address; coordinates are coarsened to a ~55km cell before they are
// written anywhere; and rows are swept after 24 hours.
//
// The live broadcaster is in-process, so this assumes a single replica. That
// already holds — the SQLite database lives on a single PVC.
type VisitTracker struct {
	q           *db.Queries
	client      *http.Client
	salt        []byte
	started     time.Time
	servingFrom string

	mu       sync.RWMutex
	ring     []VisitPing
	live     map[string]time.Time
	subs     map[chan VisitPing]struct{}
	geoMem   map[string]geoResult
	inflight map[string]struct{}

	snapMu    sync.Mutex
	snapCache StatsSnapshot
	snapAt    time.Time

	bucket *tokenBucket
}

func NewVisitTracker(q *db.Queries) *VisitTracker {
	return &VisitTracker{
		q:           q,
		client:      &http.Client{Timeout: 3 * time.Second},
		salt:        resolveVisitSalt(),
		started:     time.Now(),
		servingFrom: getEnvOr("SERVING_REGION", "EU"),
		ring:        make([]VisitPing, 0, ringCapacity),
		live:        map[string]time.Time{},
		subs:        map[chan VisitPing]struct{}{},
		geoMem:      map[string]geoResult{},
		inflight:    map[string]struct{}{},
		// ip-api.com allows 45 requests/minute on the free tier. Stay under it.
		bucket: newTokenBucket(40, time.Minute),
	}
}

// Start runs the retention sweep until ctx is cancelled.
func (t *VisitTracker) Start(ctx context.Context) {
	if t == nil || t.q == nil {
		return
	}
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		t.sweep(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				t.sweep(ctx)
			}
		}
	}()
}

func (t *VisitTracker) sweep(ctx context.Context) {
	_ = t.q.DeleteVisitsBefore(ctx, statsWindow)

	cutoff := time.Now().Add(-liveWindow)
	t.mu.Lock()
	for hash, seen := range t.live {
		if seen.Before(cutoff) {
			delete(t.live, hash)
		}
	}
	t.mu.Unlock()
}

// ShouldTrack reports whether a request is a real page view worth plotting.
// Assets, API calls, admin pages and obvious crawlers are all skipped.
func ShouldTrack(r *http.Request) bool {
	if r.Method != http.MethodGet {
		return false
	}
	if !strings.Contains(r.Header.Get("Accept"), "text/html") {
		return false
	}
	if isCrawler(r.Header.Get("User-Agent")) {
		return false
	}

	path := r.URL.Path
	for _, prefix := range []string{"/assets/", "/api/", "/templui/", "/admin"} {
		if strings.HasPrefix(path, prefix) {
			return false
		}
	}
	switch path {
	case "/healthz", "/robots.txt", "/sitemap.xml", "/favicon.ico":
		return false
	}
	return true
}

var crawlerMarkers = []string{
	"bot", "crawl", "spider", "slurp", "curl", "wget",
	"headless", "preview", "monitor", "python-requests", "go-http-client",
}

func isCrawler(userAgent string) bool {
	ua := strings.ToLower(userAgent)
	if ua == "" {
		return true
	}
	for _, marker := range crawlerMarkers {
		if strings.Contains(ua, marker) {
			return true
		}
	}
	return false
}

// Touch marks a visitor as present. It runs before the handler so that the
// page a visitor is loading already counts them — otherwise the live figure
// always trails by one page view and a lone visitor is told nobody is here.
// It costs a hash and a map write, and needs no response status.
func (t *VisitTracker) Touch(r *http.Request) {
	if t == nil {
		return
	}
	ip := clientIP(r)
	if net.ParseIP(ip) == nil {
		return
	}

	t.mu.Lock()
	t.live[t.hashIP(ip)] = time.Now()
	t.mu.Unlock()
}

// Record persists a page view and puts it on the globe. Everything that needs
// the request is read synchronously; the geo lookup and database write happen
// off the request path.
func (t *VisitTracker) Record(r *http.Request) {
	if t == nil || t.q == nil {
		return
	}

	ip := clientIP(r)
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return
	}

	go t.resolveAndStore(parsed, t.hashIP(ip), r.URL.Path)
}

func (t *VisitTracker) resolveAndStore(ip net.IP, hash, path string) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	geo, ok := t.geoFor(ctx, ip)
	if !ok {
		return
	}

	lat, lon := coarsen(geo.lat, geo.lon, hash)

	if err := t.q.RecordVisit(ctx, db.RecordVisitParams{
		VisitorHash: hash,
		CountryCode: geo.countryCode,
		City:        geo.city,
		Lat:         lat,
		Lon:         lon,
		Path:        path,
	}); err != nil {
		return
	}

	ping := VisitPing{
		CountryCode: geo.countryCode,
		City:        geo.city,
		Lat:         lat,
		Lon:         lon,
		Path:        path,
		At:          time.Now().UTC(),
	}
	t.push(ping)
	t.broadcast(ping)
	t.invalidateSnapshot()
}

// geoFor resolves an address to a city, checking the in-memory cache, then the
// database, then the upstream provider. All three are keyed by network prefix.
func (t *VisitTracker) geoFor(ctx context.Context, ip net.IP) (geoResult, bool) {
	if ip.IsLoopback() {
		// Local development: put the dot where the developer actually is.
		return geoResult{countryCode: "PT", city: "Porto", lat: 41.15, lon: -8.61}, true
	}
	if ip.IsPrivate() || ip.IsUnspecified() || ip.IsLinkLocalUnicast() {
		return geoResult{}, false
	}

	prefix := networkPrefix(ip)
	if prefix == "" {
		return geoResult{}, false
	}

	t.mu.RLock()
	cached, hit := t.geoMem[prefix]
	t.mu.RUnlock()
	if hit {
		return cached, true
	}

	if row, err := t.q.GetGeoCache(ctx, prefix); err == nil {
		result := geoResult{
			countryCode: row.CountryCode,
			city:        row.City,
			lat:         row.Lat,
			lon:         row.Lon,
		}
		t.rememberGeo(prefix, result)
		return result, true
	}

	// Only one lookup per prefix may be in flight, and only if the upstream
	// budget allows it. A dropped lookup means a dropped dot, never a stall.
	t.mu.Lock()
	if _, busy := t.inflight[prefix]; busy {
		t.mu.Unlock()
		return geoResult{}, false
	}
	t.inflight[prefix] = struct{}{}
	t.mu.Unlock()

	defer func() {
		t.mu.Lock()
		delete(t.inflight, prefix)
		t.mu.Unlock()
	}()

	if !t.bucket.allow() {
		return geoResult{}, false
	}

	result, ok := t.lookupUpstream(ctx, ip.String())
	if !ok {
		return geoResult{}, false
	}

	t.rememberGeo(prefix, result)
	_ = t.q.UpsertGeoCache(ctx, db.UpsertGeoCacheParams{
		Prefix:      prefix,
		CountryCode: result.countryCode,
		City:        result.city,
		Lat:         result.lat,
		Lon:         result.lon,
	})
	return result, true
}

func (t *VisitTracker) rememberGeo(prefix string, result geoResult) {
	t.mu.Lock()
	t.geoMem[prefix] = result
	t.mu.Unlock()
}

func (t *VisitTracker) lookupUpstream(ctx context.Context, ip string) (geoResult, bool) {
	url := "http://ip-api.com/json/" + ip + "?fields=status,countryCode,city,lat,lon"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return geoResult{}, false
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return geoResult{}, false
	}
	defer func() { _ = resp.Body.Close() }()

	var decoded struct {
		Status      string  `json:"status"`
		CountryCode string  `json:"countryCode"`
		City        string  `json:"city"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return geoResult{}, false
	}
	if decoded.Status != "success" || len(decoded.CountryCode) != 2 {
		return geoResult{}, false
	}

	city := strings.TrimSpace(decoded.City)
	if city == "" {
		city = "Unknown"
	}
	return geoResult{
		countryCode: strings.ToUpper(decoded.CountryCode),
		city:        city,
		lat:         decoded.Lat,
		lon:         decoded.Lon,
	}, true
}

func (t *VisitTracker) push(ping VisitPing) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.ring) < ringCapacity {
		t.ring = append(t.ring, ping)
		return
	}
	copy(t.ring, t.ring[1:])
	t.ring[ringCapacity-1] = ping
}

// Recent returns the newest pings held in memory, newest first.
func (t *VisitTracker) Recent(n int) []VisitPing {
	if t == nil {
		return nil
	}
	t.mu.RLock()
	defer t.mu.RUnlock()

	if n > len(t.ring) {
		n = len(t.ring)
	}
	out := make([]VisitPing, 0, n)
	for i := len(t.ring) - 1; i >= len(t.ring)-n; i-- {
		out = append(out, t.ring[i])
	}
	return out
}

// Subscribe returns a channel of live pings and a function to unsubscribe.
// Sends are non-blocking: a slow reader loses pings rather than holding up a
// page render.
func (t *VisitTracker) Subscribe() (<-chan VisitPing, func()) {
	ch := make(chan VisitPing, 16)
	if t == nil {
		close(ch)
		return ch, func() {}
	}

	t.mu.Lock()
	if len(t.subs) >= maxSubscribers {
		t.mu.Unlock()
		close(ch)
		return ch, func() {}
	}
	t.subs[ch] = struct{}{}
	t.mu.Unlock()

	var once sync.Once
	return ch, func() {
		once.Do(func() {
			t.mu.Lock()
			delete(t.subs, ch)
			t.mu.Unlock()
			close(ch)
		})
	}
}

func (t *VisitTracker) broadcast(ping VisitPing) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for ch := range t.subs {
		select {
		case ch <- ping:
		default:
		}
	}
}

func (t *VisitTracker) invalidateSnapshot() {
	t.snapMu.Lock()
	t.snapAt = time.Time{}
	t.snapMu.Unlock()
}

// Snapshot returns the 24h aggregates. It is cached briefly because the
// landing page renders it on every request.
func (t *VisitTracker) Snapshot(ctx context.Context) StatsSnapshot {
	if t == nil {
		return StatsSnapshot{}
	}

	t.snapMu.Lock()
	defer t.snapMu.Unlock()

	if time.Since(t.snapAt) < snapshotTTL {
		snap := t.snapCache
		snap.VisitorsNow = t.visitorsNow()
		snap.Uptime = time.Since(t.started)
		return snap
	}

	snap := StatsSnapshot{
		VisitorsNow: t.visitorsNow(),
		Uptime:      time.Since(t.started),
		ServingFrom: t.servingFrom,
	}

	if t.q != nil {
		if n, err := t.q.CountVisitsSince(ctx, statsWindow); err == nil {
			snap.Views24h = int(n)
		}
		if n, err := t.q.CountVisitorsSince(ctx, statsWindow); err == nil {
			snap.Visitors24h = int(n)
		}
		if n, err := t.q.CountCountriesSince(ctx, statsWindow); err == nil {
			snap.Countries24h = int(n)
		}
		if rows, err := t.q.TopCountriesSince(ctx, db.TopCountriesSinceParams{
			Datetime: statsWindow,
			Limit:    12,
		}); err == nil {
			for _, row := range rows {
				snap.TopCountries = append(snap.TopCountries, CountryCount{
					Code:   row.CountryCode,
					Flag:   FlagEmoji(row.CountryCode),
					Visits: int(row.Visits),
				})
			}
		}
		if rows, err := t.q.TopPathsSince(ctx, db.TopPathsSinceParams{
			Datetime: statsWindow,
			Limit:    6,
		}); err == nil {
			for _, row := range rows {
				snap.TopPaths = append(snap.TopPaths, PathCount{
					Path:   row.Path,
					Visits: int(row.Visits),
				})
			}
		}
	}

	t.snapCache = snap
	t.snapAt = time.Now()
	return snap
}

// VisitsInWindow returns every visit still inside the 24h window, newest
// first. This is the initial payload the globe draws before the stream opens.
func (t *VisitTracker) VisitsInWindow(ctx context.Context, limit int) []VisitPing {
	if t == nil || t.q == nil {
		return nil
	}
	rows, err := t.q.ListRecentVisits(ctx, db.ListRecentVisitsParams{
		Datetime: statsWindow,
		Limit:    int64(limit),
	})
	if err != nil {
		return nil
	}

	out := make([]VisitPing, 0, len(rows))
	for _, row := range rows {
		at, err := time.Parse("2006-01-02 15:04:05", row.CreatedAt)
		if err != nil {
			at = time.Now().UTC()
		}
		out = append(out, VisitPing{
			CountryCode: row.CountryCode,
			City:        row.City,
			Lat:         row.Lat,
			Lon:         row.Lon,
			Path:        row.Path,
			At:          at.UTC(),
		})
	}
	return out
}

func (t *VisitTracker) visitorsNow() int {
	cutoff := time.Now().Add(-liveWindow)
	t.mu.RLock()
	defer t.mu.RUnlock()

	count := 0
	for _, seen := range t.live {
		if seen.After(cutoff) {
			count++
		}
	}
	return count
}

func (t *VisitTracker) hashIP(ip string) string {
	return hashWithSalt(t.salt, ip)
}

// hashWithSalt is the one place an address turns into an identifier. Rotating
// the salt unlinks every stored hash from every future one, which is why the
// fallback below is random per process.
func hashWithSalt(salt []byte, value string) string {
	sum := sha256.Sum256(append(append([]byte{}, salt...), []byte(value)...))
	return hex.EncodeToString(sum[:16])
}

// resolveVisitSalt finds the salt in the first place that has one: an explicit
// env var, then a file on the durable volume, then a fresh random value which
// it writes to that file for next time.
//
// Keeping it on the volume rather than in a Secret means it is never in git and
// never in a config store, it survives redeploys so unique-visitor counts stay
// stable, and rotating it is deleting one file. If nothing can be persisted the
// process still runs with an in-memory salt — hashes then stop being comparable
// across restarts, which costs accuracy but never privacy.
func resolveVisitSalt() []byte {
	if explicit := os.Getenv("VISIT_SALT"); explicit != "" {
		return []byte(explicit)
	}

	path := resolveVisitSaltPath()
	if stored, err := os.ReadFile(path); err == nil {
		if trimmed := strings.TrimSpace(string(stored)); trimmed != "" {
			return []byte(trimmed)
		}
	}

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		// A time-derived salt still beats hashing unsalted.
		return []byte(time.Now().Format(time.RFC3339Nano))
	}
	encoded := hex.EncodeToString(salt)

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err == nil {
		// 0600: readable only by the process that owns it.
		_ = os.WriteFile(path, []byte(encoded), 0o600)
	}
	return []byte(encoded)
}

func resolveVisitSaltPath() string {
	if explicit := os.Getenv("VISIT_SALT_PATH"); explicit != "" {
		return explicit
	}
	if dbPath := os.Getenv("DATABASE_PATH"); dbPath != "" {
		return filepath.Join(filepath.Dir(dbPath), "visit-salt")
	}
	return "visit-salt"
}

// networkPrefix reduces an address to the block geo lookups are cached against:
// /24 for IPv4, /48 for IPv6. The host part is dropped and never stored.
func networkPrefix(ip net.IP) string {
	if v4 := ip.To4(); v4 != nil {
		return net.IP(append(append([]byte{}, v4[:3]...), 0)).String() + "/24"
	}
	if v6 := ip.To16(); v6 != nil {
		masked := make([]byte, 16)
		copy(masked, v6[:6])
		return net.IP(masked).String() + "/48"
	}
	return ""
}

// coarsen snaps a coordinate to a half-degree cell (~55km) and then scatters it
// deterministically inside that cell, so a dot can never be traced back to an
// address and repeat visitors do not stack on one pixel.
func coarsen(lat, lon float64, hash string) (float64, float64) {
	const cell = 0.5

	snappedLat := math.Round(lat/cell) * cell
	snappedLon := math.Round(lon/cell) * cell

	jitterLat, jitterLon := jitterFor(hash)
	outLat := snappedLat + jitterLat*cell
	outLon := snappedLon + jitterLon*cell

	return math.Max(-85, math.Min(85, outLat)), math.Mod(outLon+540, 360) - 180
}

// jitterFor derives two values in [-0.5, 0.5) from a visitor hash.
func jitterFor(hash string) (float64, float64) {
	sum := sha256.Sum256([]byte(hash))
	a := binary.BigEndian.Uint32(sum[0:4])
	b := binary.BigEndian.Uint32(sum[4:8])
	return float64(a)/float64(math.MaxUint32) - 0.5, float64(b)/float64(math.MaxUint32) - 0.5
}

// FlagEmoji turns an ISO country code into its regional-indicator flag.
func FlagEmoji(code string) string {
	code = strings.ToUpper(strings.TrimSpace(code))
	if len(code) != 2 {
		return ""
	}
	runes := make([]rune, 0, 2)
	for _, c := range code {
		if c < 'A' || c > 'Z' {
			return ""
		}
		runes = append(runes, rune(0x1F1E6+(c-'A')))
	}
	return string(runes)
}

// tokenBucket caps how fast we call the upstream geo provider.
type tokenBucket struct {
	mu     sync.Mutex
	tokens float64
	burst  float64
	perSec float64
	last   time.Time
}

func newTokenBucket(burst int, per time.Duration) *tokenBucket {
	return &tokenBucket{
		tokens: float64(burst),
		burst:  float64(burst),
		perSec: float64(burst) / per.Seconds(),
		last:   time.Now(),
	}
}

func (b *tokenBucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	b.tokens = math.Min(b.burst, b.tokens+now.Sub(b.last).Seconds()*b.perSec)
	b.last = now

	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

func getEnvOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

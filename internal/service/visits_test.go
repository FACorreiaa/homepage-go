package service

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"myapp/internal/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func htmlRequest(method, path, userAgent string) *http.Request {
	r := httptest.NewRequest(method, path, nil)
	r.Header.Set("Accept", "text/html,application/xhtml+xml")
	r.Header.Set("User-Agent", userAgent)
	return r
}

const browserUA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 Chrome/140 Safari/537.36"

func TestShouldTrack(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{"page view", htmlRequest("GET", "/projects", browserUA), true},
		{"home", htmlRequest("GET", "/", browserUA), true},
		{"post is not a view", htmlRequest("POST", "/proposal", browserUA), false},
		{"asset", htmlRequest("GET", "/assets/css/output.css", browserUA), false},
		{"api", htmlRequest("GET", "/api/stats", browserUA), false},
		{"stats stream", htmlRequest("GET", "/api/stats/stream", browserUA), false},
		{"templui script", htmlRequest("GET", "/templui/js/x.js", browserUA), false},
		{"admin", htmlRequest("GET", "/admin/dashboard", browserUA), false},
		{"health check", htmlRequest("GET", "/healthz", browserUA), false},
		{"robots", htmlRequest("GET", "/robots.txt", browserUA), false},
		{"sitemap", htmlRequest("GET", "/sitemap.xml", browserUA), false},
		{"crawler", htmlRequest("GET", "/", "Googlebot/2.1"), false},
		{"empty user agent", htmlRequest("GET", "/", ""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ShouldTrack(tt.req))
		})
	}
}

func TestShouldTrackRequiresHTML(t *testing.T) {
	r := httptest.NewRequest("GET", "/projects", nil)
	r.Header.Set("User-Agent", browserUA)
	r.Header.Set("Accept", "application/json")

	assert.False(t, ShouldTrack(r), "a JSON fetch is not a page view")
}

func TestNetworkPrefixDropsHostPart(t *testing.T) {
	tests := []struct {
		ip   string
		want string
	}{
		{"203.0.113.42", "203.0.113.0/24"},
		{"203.0.113.199", "203.0.113.0/24"},
		{"8.8.8.8", "8.8.8.0/24"},
		{"2001:db8:1234:5678::1", "2001:db8:1234::/48"},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			assert.Equal(t, tt.want, networkPrefix(net.ParseIP(tt.ip)))
		})
	}
}

func TestNetworkPrefixSharedAcrossHostsInBlock(t *testing.T) {
	a := networkPrefix(net.ParseIP("198.51.100.7"))
	b := networkPrefix(net.ParseIP("198.51.100.240"))

	assert.Equal(t, a, b, "two hosts in one /24 must share a cache key")
}

func TestCoarsenStaysNearOriginAndIsStable(t *testing.T) {
	const lat, lon = 41.1579, -8.6291 // Porto
	hash := hashWithSalt([]byte("salt"), "203.0.113.42")

	gotLat, gotLon := coarsen(lat, lon, hash)
	againLat, againLon := coarsen(lat, lon, hash)

	assert.Equal(t, gotLat, againLat, "coarsening must be deterministic")
	assert.Equal(t, gotLon, againLon)
	assert.InDelta(t, lat, gotLat, 0.75, "dot must stay in the right region")
	assert.InDelta(t, lon, gotLon, 0.75)
	assert.NotEqual(t, lat, gotLat, "exact coordinates must not survive")
}

func TestCoarsenSeparatesDifferentVisitorsInSameCity(t *testing.T) {
	const lat, lon = 51.5074, -0.1278 // London

	a1, a2 := coarsen(lat, lon, hashWithSalt([]byte("salt"), "203.0.113.1"))
	b1, b2 := coarsen(lat, lon, hashWithSalt([]byte("salt"), "203.0.113.2"))

	assert.False(t, a1 == b1 && a2 == b2, "visitors should not stack on one pixel")
}

func TestCoarsenClampsToValidRange(t *testing.T) {
	lat, lon := coarsen(89.9, 179.9, hashWithSalt([]byte("salt"), "203.0.113.9"))

	assert.LessOrEqual(t, lat, 85.0)
	assert.GreaterOrEqual(t, lat, -85.0)
	assert.Less(t, lon, 180.0)
	assert.GreaterOrEqual(t, lon, -180.0)
}

func TestHashWithSaltIsStableAndSaltDependent(t *testing.T) {
	a := hashWithSalt([]byte("salt-one"), "203.0.113.42")
	again := hashWithSalt([]byte("salt-one"), "203.0.113.42")
	rotated := hashWithSalt([]byte("salt-two"), "203.0.113.42")

	assert.Equal(t, a, again)
	assert.NotEqual(t, a, rotated, "rotating the salt must unlink stored hashes")
	assert.NotContains(t, a, "203.0.113", "the address must not survive hashing")
}

func TestFlagEmoji(t *testing.T) {
	assert.Equal(t, "🇵🇹", FlagEmoji("PT"))
	assert.Equal(t, "🇩🇪", FlagEmoji("de"))
	assert.Equal(t, "", FlagEmoji("XYZ"))
	assert.Equal(t, "", FlagEmoji(""))
	assert.Equal(t, "", FlagEmoji("1A"))
}

func TestTokenBucketCapsUpstreamCalls(t *testing.T) {
	bucket := newTokenBucket(3, time.Minute)

	assert.True(t, bucket.allow())
	assert.True(t, bucket.allow())
	assert.True(t, bucket.allow())
	assert.False(t, bucket.allow(), "budget is spent")
}

func TestRingBufferKeepsNewestAndNeverGrows(t *testing.T) {
	tracker := NewVisitTracker(nil)

	for i := 0; i < ringCapacity+50; i++ {
		tracker.push(VisitPing{City: "city", Path: "/", Lat: float64(i)})
	}

	require.Len(t, tracker.ring, ringCapacity)

	recent := tracker.Recent(3)
	require.Len(t, recent, 3)
	assert.Equal(t, float64(ringCapacity+49), recent[0].Lat, "newest first")
	assert.Equal(t, float64(ringCapacity+48), recent[1].Lat)
}

func TestRecentHandlesEmptyAndOversizedRequests(t *testing.T) {
	tracker := NewVisitTracker(nil)
	assert.Empty(t, tracker.Recent(10))

	tracker.push(VisitPing{City: "Porto"})
	assert.Len(t, tracker.Recent(100), 1)
}

func TestSubscribeReceivesBroadcast(t *testing.T) {
	tracker := NewVisitTracker(nil)
	ch, unsubscribe := tracker.Subscribe()
	defer unsubscribe()

	tracker.broadcast(VisitPing{City: "Porto", CountryCode: "PT"})

	select {
	case got := <-ch:
		assert.Equal(t, "Porto", got.City)
	case <-time.After(time.Second):
		t.Fatal("subscriber never received the ping")
	}
}

func TestBroadcastDoesNotBlockOnSlowSubscriber(t *testing.T) {
	tracker := NewVisitTracker(nil)
	_, unsubscribe := tracker.Subscribe()
	defer unsubscribe()

	done := make(chan struct{})
	go func() {
		// Far more pings than the subscriber buffer holds. A slow reader must
		// lose pings rather than stall the caller.
		for i := 0; i < 500; i++ {
			tracker.broadcast(VisitPing{City: "Porto"})
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("broadcast blocked on a subscriber that never reads")
	}
}

func TestUnsubscribeIsIdempotent(t *testing.T) {
	tracker := NewVisitTracker(nil)
	_, unsubscribe := tracker.Subscribe()

	unsubscribe()
	assert.NotPanics(t, unsubscribe, "double unsubscribe must not close a closed channel")

	tracker.mu.RLock()
	defer tracker.mu.RUnlock()
	assert.Empty(t, tracker.subs)
}

func TestSubscriberCapIsEnforced(t *testing.T) {
	tracker := NewVisitTracker(nil)

	for i := 0; i < maxSubscribers; i++ {
		_, unsubscribe := tracker.Subscribe()
		defer unsubscribe()
	}

	ch, unsubscribe := tracker.Subscribe()
	defer unsubscribe()

	_, open := <-ch
	assert.False(t, open, "an over-cap subscriber gets a closed channel, not a slot")
}

// trackerWithDB builds a tracker over a real in-memory database, so Record
// runs its full path instead of short-circuiting on a nil queries handle.
func trackerWithDB(t *testing.T) (*VisitTracker, *sql.DB) {
	t.Helper()

	conn, err := db.Open(":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, conn.Close()) })

	return NewVisitTracker(db.New(conn)), conn
}

func TestRecordPlotsNothingForPrivateAddresses(t *testing.T) {
	tracker, _ := trackerWithDB(t)

	for _, ip := range []string{"10.0.0.4", "192.168.1.9", "169.254.4.4"} {
		r := htmlRequest("GET", "/", browserUA)
		r.Header.Set("X-Forwarded-For", ip)
		tracker.Touch(r)
		tracker.Record(r)
	}
	waitForVisits(t, tracker, 3)

	// A LAN address cannot be placed on the globe and must never reach the geo
	// provider — but it is still a real person reading a real page.
	assert.Equal(t, 3, tracker.visitorsNow())
	assert.Empty(t, tracker.Recent(10), "unlocatable addresses must not produce dots")
	assert.Empty(t, tracker.VisitsInWindow(t.Context(), 10), "and must not reach the globe payload")

	snap := tracker.Snapshot(t.Context())
	assert.Equal(t, 3, snap.Views24h, "an unplaceable visit is still a page view")
	assert.Equal(t, 3, snap.Visitors24h)
	assert.Zero(t, snap.Countries24h, "but it belongs to no country")
	assert.Empty(t, snap.TopCountries, "and must never show a blank flag")
	require.Len(t, snap.TopPaths, 1, "the path is known even when the place is not")
	assert.Equal(t, "/", snap.TopPaths[0].Path)
	assert.Equal(t, 3, snap.TopPaths[0].Visits)
}

// A geo outage must cost the dot and nothing else. Before this, resolveAndStore
// returned before the insert, so an unreachable provider silently zeroed every
// figure on /stats while VISITORS NOW kept working — which is exactly how it
// failed in production.
func TestLocatedAndUnlocatedVisitsCoexist(t *testing.T) {
	tracker, _ := trackerWithDB(t)

	// Loopback resolves to Porto without touching the network; a LAN address
	// cannot resolve at all.
	for _, ip := range []string{"127.0.0.1", "10.0.0.4"} {
		r := htmlRequest("GET", "/projects", browserUA)
		r.Header.Set("X-Forwarded-For", ip)
		tracker.Record(r)
	}
	waitForVisits(t, tracker, 2)

	snap := tracker.Snapshot(t.Context())
	assert.Equal(t, 2, snap.Views24h, "both are page views")
	assert.Equal(t, 1, snap.Countries24h, "only one of them has a country")
	require.Len(t, snap.TopCountries, 1)
	assert.Equal(t, "PT", snap.TopCountries[0].Code)

	dots := tracker.VisitsInWindow(t.Context(), 10)
	require.Len(t, dots, 1, "only the located visit is a dot")
	assert.Equal(t, "PT", dots[0].CountryCode)
}

// waitForVisits blocks until the expected number of rows have been written.
// Record hands off to a goroutine, so the count is not there on return.
func waitForVisits(t *testing.T, tracker *VisitTracker, want int) {
	t.Helper()
	for range 100 {
		if tracker.countVisits(t.Context()) >= want {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for %d visits, saw %d", want, tracker.countVisits(t.Context()))
}

func (t *VisitTracker) countVisits(ctx context.Context) int {
	n, err := t.q.CountVisitsSince(ctx, statsWindow)
	if err != nil {
		return 0
	}
	return int(n)
}

// The live figure has to include the person loading the page that displays it,
// which is why the middleware touches before the handler runs.
func TestTouchCountsTheVisitorImmediately(t *testing.T) {
	tracker, _ := trackerWithDB(t)

	r := htmlRequest("GET", "/", browserUA)
	r.Header.Set("X-Forwarded-For", "203.0.113.77")

	assert.Zero(t, tracker.visitorsNow())
	tracker.Touch(r)
	assert.Equal(t, 1, tracker.Snapshot(t.Context()).VisitorsNow, "a lone visitor must not be told nobody is here")
}

func TestTouchIsIdempotentForOneVisitor(t *testing.T) {
	tracker, _ := trackerWithDB(t)

	for _, path := range []string{"/", "/projects", "/stats"} {
		r := htmlRequest("GET", path, browserUA)
		r.Header.Set("X-Forwarded-For", "203.0.113.77")
		tracker.Touch(r)
	}

	assert.Equal(t, 1, tracker.visitorsNow(), "one person browsing three pages is one visitor")
}

func TestRecordIgnoresUnparseableAddresses(t *testing.T) {
	tracker, _ := trackerWithDB(t)

	r := htmlRequest("GET", "/", browserUA)
	r.Header.Set("X-Forwarded-For", "not-an-ip")
	tracker.Record(r)

	assert.Zero(t, tracker.visitorsNow())
	assert.Empty(t, tracker.Recent(10))
}

func TestRecordPlacesLoopbackAndStoresNoAddress(t *testing.T) {
	tracker, conn := trackerWithDB(t)

	r := htmlRequest("GET", "/projects", browserUA)
	r.Header.Set("X-Forwarded-For", "127.0.0.1")
	tracker.Record(r)

	require.Eventually(t, func() bool { return len(tracker.Recent(1)) == 1 }, 2*time.Second, 10*time.Millisecond)

	ping := tracker.Recent(1)[0]
	assert.Equal(t, "PT", ping.CountryCode)
	assert.Equal(t, "/projects", ping.Path)

	stored := tracker.VisitsInWindow(t.Context(), 10)
	require.Len(t, stored, 1)
	assert.Equal(t, "/projects", stored[0].Path)

	// Nothing resembling the address may reach the database.
	var blob string
	require.NoError(t, conn.QueryRowContext(t.Context(),
		`SELECT group_concat(visitor_hash || '|' || country_code || '|' || city || '|' || path) FROM visits`,
	).Scan(&blob))
	assert.NotContains(t, blob, "127.0.0.1")
}

func TestSweepDropsRowsOutsideTheWindow(t *testing.T) {
	tracker, conn := trackerWithDB(t)

	require.NoError(t, tracker.q.RecordVisit(t.Context(), db.RecordVisitParams{
		VisitorHash: "hash", CountryCode: "PT", City: "Porto", Lat: 41, Lon: -8, Path: "/",
	}))
	require.Len(t, tracker.VisitsInWindow(t.Context(), 10), 1)

	// Age the row past the retention window, then sweep.
	_, err := conn.ExecContext(t.Context(), `UPDATE visits SET created_at = datetime('now', '-25 hours')`)
	require.NoError(t, err)

	tracker.sweep(t.Context())

	assert.Empty(t, tracker.VisitsInWindow(t.Context(), 10), "rows older than 24h must be swept")
}

func TestGeoForSkipsPrivateRanges(t *testing.T) {
	tracker := NewVisitTracker(nil)

	_, ok := tracker.geoFor(t.Context(), net.ParseIP("192.168.0.1"))
	assert.False(t, ok, "private addresses must never reach the upstream provider")
}

func TestVisitorsNowCountsOnlyTheLiveWindow(t *testing.T) {
	tracker := NewVisitTracker(nil)

	tracker.live["fresh"] = time.Now()
	tracker.live["stale"] = time.Now().Add(-liveWindow - time.Minute)

	assert.Equal(t, 1, tracker.visitorsNow())
}

func TestSnapshotWithoutDatabaseStillReportsOps(t *testing.T) {
	tracker := NewVisitTracker(nil)
	tracker.live["someone"] = time.Now()

	snap := tracker.Snapshot(t.Context())

	assert.Equal(t, 1, snap.VisitorsNow)
	assert.Equal(t, "EU", snap.ServingFrom)
	assert.Positive(t, snap.Uptime)
	assert.Zero(t, snap.Views24h)
}

// The two providers disagree about field names — country_code vs countryCode,
// city vs cityName — so the parsers are pinned against real response bodies.
// Getting this wrong looks exactly like a provider being down.
func TestGeoProviderParsers(t *testing.T) {
	bodies := map[string]string{
		"ipwho.is": `{"ip":"177.71.207.1","success":true,"country":"Brazil",
			"country_code":"BR","region":"Sao Paulo","city":"Sao Paulo",
			"latitude":-23.5558,"longitude":-46.6396}`,
		"freeipapi.com": `{"ipVersion":4,"ipAddress":"177.71.207.1","latitude":-23.5558,
			"longitude":-46.6396,"countryName":"Brazil","countryCode":"BR","cityName":"Sao Paulo"}`,
	}

	for _, provider := range geoProviders {
		t.Run(provider.name, func(t *testing.T) {
			body, ok := bodies[provider.name]
			require.True(t, ok, "no sample body for %s", provider.name)

			got, ok := provider.parse([]byte(body))
			require.True(t, ok, "a good response must parse")
			assert.Equal(t, "BR", got.countryCode)
			assert.Equal(t, "Sao Paulo", got.city)
			assert.InDelta(t, -23.5558, got.lat, 0.0001)
			assert.InDelta(t, -46.6396, got.lon, 0.0001)

			assert.Contains(t, provider.url("1.2.3.4"), "1.2.3.4")
			assert.True(t, strings.HasPrefix(provider.url("1.2.3.4"), "https://"),
				"plain HTTP is what broke this: ip-api's free tier is HTTP-only "+
					"and its HTTPS endpoint answers 403 without a paid key")
		})
	}
}

// A refusal must not be mistaken for a location.
func TestGeoProviderParsersRejectFailures(t *testing.T) {
	for _, tc := range []struct{ name, body string }{
		{"ipwho.is refusal", `{"ip":"1.2.3.4","success":false,"message":"Reserved range"}`},
		{"ip-api style refusal", `{"status":"fail"}`},
		{"empty object", `{}`},
		{"not json", `<html>rate limited</html>`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, provider := range geoProviders {
				_, ok := provider.parse([]byte(tc.body))
				assert.False(t, ok, "%s must reject %s", provider.name, tc.name)
			}
		})
	}
}

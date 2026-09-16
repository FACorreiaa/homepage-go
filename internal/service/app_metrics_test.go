package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func metricsServer(t *testing.T, secret string, users int64, hits *int32) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits != nil {
			atomic.AddInt32(hits, 1)
		}
		if r.Header.Get("Authorization") != "Bearer "+secret {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"users":` + itoa(users) + `}`))
	}))
	t.Cleanup(srv.Close)
	return srv
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	neg := n < 0
	if neg {
		n = -n
	}
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	if neg {
		b = append([]byte{'-'}, b...)
	}
	return string(b)
}

func TestFetchAllSuccessAndOrder(t *testing.T) {
	a := metricsServer(t, "s3cret", 42, nil)
	b := metricsServer(t, "s3cret", 0, nil)

	svc := NewAppMetricsServiceWith([]AppEntry{
		{Name: "A", Slug: "a", URL: a.URL},
		{Name: "B", Slug: "b", URL: b.URL},
	}, "s3cret", nil)

	results := svc.FetchAll(context.Background())
	require.Len(t, results, 2)
	assert.Equal(t, "a", results[0].App.Slug)
	assert.Equal(t, "b", results[1].App.Slug)
	require.NotNil(t, results[0].Users)
	assert.EqualValues(t, 42, *results[0].Users)
	// Zero is a real answer, not a failure.
	require.NotNil(t, results[1].Users)
	assert.EqualValues(t, 0, *results[1].Users)
	assert.EqualValues(t, 42, TotalUsers(results))
	assert.Equal(t, 2, ReportingCount(results))
}

func TestFetchAllWrongSecretIsUnavailable(t *testing.T) {
	srv := metricsServer(t, "right", 7, nil)
	svc := NewAppMetricsServiceWith([]AppEntry{{Slug: "a", URL: srv.URL}}, "wrong", nil)

	results := svc.FetchAll(context.Background())
	require.Len(t, results, 1)
	assert.Nil(t, results[0].Users)
	assert.ErrorContains(t, results[0].Err, "status 401")
}

func TestFetchAllUnconfiguredURL(t *testing.T) {
	svc := NewAppMetricsServiceWith([]AppEntry{{Slug: "a"}}, "s", nil)
	results := svc.FetchAll(context.Background())
	require.Len(t, results, 1)
	assert.Nil(t, results[0].Users)
	assert.ErrorIs(t, results[0].Err, ErrNotConfigured)
}

func TestFetchAllTimeoutDoesNotBlockOthers(t *testing.T) {
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	t.Cleanup(slow.Close)
	fast := metricsServer(t, "s", 3, nil)

	svc := NewAppMetricsServiceWith([]AppEntry{
		{Slug: "slow", URL: slow.URL},
		{Slug: "fast", URL: fast.URL},
	}, "s", &http.Client{Timeout: 200 * time.Millisecond})

	start := time.Now()
	results := svc.FetchAll(context.Background())
	assert.Less(t, time.Since(start), 2*time.Second)
	assert.Nil(t, results[0].Users)
	assert.Error(t, results[0].Err)
	require.NotNil(t, results[1].Users)
	assert.EqualValues(t, 3, *results[1].Users)
}

func TestFetchAllMalformedBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"total":5}`))
	}))
	t.Cleanup(srv.Close)
	svc := NewAppMetricsServiceWith([]AppEntry{{Slug: "a", URL: srv.URL}}, "s", nil)
	results := svc.FetchAll(context.Background())
	assert.Nil(t, results[0].Users)
	assert.ErrorContains(t, results[0].Err, "missing users")
}

func TestFetchAllCachesWithinTTL(t *testing.T) {
	var hits int32
	srv := metricsServer(t, "s", 1, &hits)
	svc := NewAppMetricsServiceWith([]AppEntry{{Slug: "a", URL: srv.URL}}, "s", nil)

	svc.FetchAll(context.Background())
	svc.FetchAll(context.Background())
	assert.EqualValues(t, 1, atomic.LoadInt32(&hits), "second call within TTL must be served from cache")

	svc.SetTTL(0)
	svc.FetchAll(context.Background())
	svc.FetchAll(context.Background())
	assert.EqualValues(t, 3, atomic.LoadInt32(&hits), "TTL 0 disables the cache")
}

func TestNilServiceIsEmptyBoard(t *testing.T) {
	var svc *AppMetricsService
	assert.Nil(t, svc.FetchAll(context.Background()))
	assert.Nil(t, svc.Apps())
}

func TestAppsReturnsRegistryUnknown(t *testing.T) {
	svc := NewAppMetricsServiceWith([]AppEntry{{Slug: "a"}, {Slug: "b"}}, "", nil)
	rows := svc.Apps()
	require.Len(t, rows, 2)
	for _, r := range rows {
		assert.Nil(t, r.Users)
		assert.ErrorIs(t, r.Err, ErrNotConfigured)
	}
}

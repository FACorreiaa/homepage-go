package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"myapp/internal/handler"
	"myapp/internal/service"
	"myapp/ui/pages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The globe reads these four attributes to update the figures live. Restyling
// the page must not drop them, so they are asserted rather than eyeballed.
var statsMetricAttrs = []string{
	`data-metric="visitorsNow"`,
	`data-metric="views24h"`,
	`data-metric="visitors24h"`,
	`data-metric="countries24h"`,
}

// A nil tracker is a legal zero-traffic tracker: Snapshot guards its receiver.
// That makes the empty-state path reachable without a database.
func TestStatsPageEmptyState(t *testing.T) {
	h := &handler.StatsHandler{}
	r := httptest.NewRequest("GET", "/stats", nil)
	w := httptest.NewRecorder()

	h.Page(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	for _, attr := range statsMetricAttrs {
		assert.Contains(t, body, attr)
	}
	assert.Contains(t, body, `data-globe`)
	assert.Contains(t, body, `/assets/static/globe.js`)
	// Both lists are empty here; the proportion bars must not divide by zero.
	assert.Contains(t, body, "Nobody yet today.")
	assert.Contains(t, body, "No page views recorded in the window.")
	assert.NotContains(t, body, "stat-row-bar")
}

func TestStatsPagePopulated(t *testing.T) {
	snapshot := service.StatsSnapshot{
		VisitorsNow:  3,
		Views24h:     412,
		Visitors24h:  97,
		Countries24h: 11,
		TopCountries: []service.CountryCount{
			{Code: "PT", Flag: "🇵🇹", Visits: 40},
			{Code: "DE", Flag: "🇩🇪", Visits: 10},
		},
		TopPaths: []service.PathCount{
			{Path: "/projects", Visits: 80},
			{Path: "/about", Visits: 20},
		},
		Uptime:      90 * time.Minute,
		ServingFrom: "EU",
	}

	var sb strings.Builder
	require.NoError(t, pages.Stats(snapshot).Render(context.Background(), &sb))
	body := sb.String()

	for _, attr := range statsMetricAttrs {
		assert.Contains(t, body, attr)
	}
	assert.Contains(t, body, `data-live="true"`)
	assert.Contains(t, body, "PT")
	assert.Contains(t, body, `href="/projects"`)
	// Live figure, so the globe has something to overwrite.
	assert.Contains(t, body, ">3<")

	// Bars are scaled against the biggest row in their own list, and the style
	// attribute must survive templ verbatim — it is emitted via templ.Raw.
	assert.Contains(t, body, `<span class="stat-row-bar" style="width:100%"`)
	assert.Contains(t, body, `<span class="stat-row-bar" style="width:25%"`)
}

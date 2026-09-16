package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"myapp/internal/handler"
	"myapp/internal/service"
	"myapp/ui/pages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A nil service is the "nothing wired up" board: still a 200, still a page.
func TestAppsPageWithNoService(t *testing.T) {
	h := &handler.AppsHandler{}
	r := httptest.NewRequest("GET", "/apps", nil)
	w := httptest.NewRecorder()

	h.Page(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "public, max-age=60", w.Header().Get("Cache-Control"))
	body := w.Body.String()
	assert.Contains(t, body, `data-metric="apps">0<`)
	assert.Contains(t, body, "No apps registered.")
	assert.Contains(t, body, `<meta name="robots" content="noindex`)
}

func TestAppsPageUnconfiguredRegistry(t *testing.T) {
	svc := service.NewAppMetricsServiceWith([]service.AppEntry{
		{Name: "Norviq", Slug: "norviq", Description: "d"},
		{Name: "Loci", Slug: "loci", Description: "d"},
	}, "", nil)
	h := &handler.AppsHandler{Metrics: svc}
	w := httptest.NewRecorder()

	h.Page(w, httptest.NewRequest("GET", "/apps", nil))

	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `data-app="norviq"`)
	assert.Contains(t, body, `data-app="loci"`)
	assert.Equal(t, 2, strings.Count(body, `data-state="unconfigured"`))
	assert.Contains(t, body, "not connected yet")
	assert.Contains(t, body, `data-metric="reporting">0<`)
	assert.Contains(t, body, `data-metric="usersTotal">0<`)
}

func TestAppsPagePopulated(t *testing.T) {
	n := int64(12345)
	zero := int64(0)
	results := []service.AppMetricsResult{
		{App: service.AppEntry{Name: "Norviq", Slug: "norviq", Description: "Stock plans."}, Users: &n},
		{App: service.AppEntry{Name: "Loci", Slug: "loci", Description: "Travel."}, Users: &zero},
		{App: service.AppEntry{Name: "North", Slug: "north", Description: "Health."}, Err: errors.New("status 502")},
		{App: service.AppEntry{Name: "Seshat", Slug: "seshat", Description: "Lessons."}, Err: service.ErrNotConfigured},
	}

	var sb strings.Builder
	require.NoError(t, pages.Apps(results).Render(context.Background(), &sb))
	body := sb.String()

	// Thousands are grouped with a space; zero is a real figure, not a dash.
	assert.Contains(t, body, `data-live="true">12 345<`)
	assert.Contains(t, body, `data-app="loci" data-state="ok"`)
	assert.Contains(t, body, `data-live="true">0<`)
	// A failed fetch and an unconfigured app are told apart on the card.
	assert.Contains(t, body, `data-app="north" data-state="unavailable"`)
	assert.Contains(t, body, "unavailable right now")
	assert.Contains(t, body, `data-app="seshat" data-state="unconfigured"`)
	assert.Contains(t, body, "not connected yet")
	assert.Equal(t, 2, strings.Count(body, "&mdash;"))

	assert.Contains(t, body, `data-metric="apps">4<`)
	assert.Contains(t, body, `data-metric="reporting">2<`)
	assert.Contains(t, body, `data-metric="usersTotal">12 345<`)
}

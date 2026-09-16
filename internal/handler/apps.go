package handler

import (
	"net/http"

	"myapp/internal/service"
	"myapp/ui/pages"
)

// AppsHandler serves /apps: one card per product with its total user count.
type AppsHandler struct {
	// Metrics may be nil, in which case the board renders every app as
	// unavailable. That keeps the page reachable in tests and in a local run
	// with nothing configured.
	Metrics *service.AppMetricsService
}

func (h *AppsHandler) Page(w http.ResponseWriter, r *http.Request) {
	results := h.Metrics.FetchAll(r.Context())
	// The counts are a minute stale by design (see AppMetricsService), so a
	// short shared cache in front is consistent with what the page shows.
	w.Header().Set("Cache-Control", "public, max-age=60")
	if err := pages.Apps(results).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

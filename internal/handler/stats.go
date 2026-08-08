package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"myapp/internal/service"
	"myapp/ui/pages"
)

// heartbeatInterval keeps the SSE connection alive through proxy idle
// timeouts on a site that can go a long while between visitors.
const heartbeatInterval = 25 * time.Second

type StatsHandler struct {
	Tracker *service.VisitTracker
}

func (h *StatsHandler) Page(w http.ResponseWriter, r *http.Request) {
	snapshot := h.Tracker.Snapshot(r.Context())
	if err := pages.Stats(snapshot).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

// JSON is the globe's initial payload: every visit still inside the window,
// plus the same aggregates the page renders server-side.
func (h *StatsHandler) JSON(w http.ResponseWriter, r *http.Request) {
	snapshot := h.Tracker.Snapshot(r.Context())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	payload := struct {
		Visits       []service.VisitPing `json:"visits"`
		VisitorsNow  int                 `json:"visitorsNow"`
		Views24h     int                 `json:"views24h"`
		Visitors24h  int                 `json:"visitors24h"`
		Countries24h int                 `json:"countries24h"`
	}{
		Visits:       h.Tracker.VisitsInWindow(r.Context(), 2000),
		VisitorsNow:  snapshot.VisitorsNow,
		Views24h:     snapshot.Views24h,
		Visitors24h:  snapshot.Visitors24h,
		Countries24h: snapshot.Countries24h,
	}
	if payload.Visits == nil {
		payload.Visits = []service.VisitPing{}
	}

	_ = json.NewEncoder(w).Encode(payload)
}

// Stream pushes each new visit to the globe as it happens.
func (h *StatsHandler) Stream(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// Tell any buffering proxy in front of us to pass bytes straight through.
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)

	pings, unsubscribe := h.Tracker.Subscribe()
	defer unsubscribe()

	_, _ = fmt.Fprint(w, ": connected\n\n")
	flusher.Flush()

	heartbeat := time.NewTicker(heartbeatInterval)
	defer heartbeat.Stop()

	for {
		select {
		case <-r.Context().Done():
			return

		case ping, open := <-pings:
			if !open {
				return
			}
			encoded, err := json.Marshal(ping)
			if err != nil {
				continue
			}
			_, _ = fmt.Fprintf(w, "event: visit\ndata: %s\n\n", encoded)
			flusher.Flush()

		case <-heartbeat.C:
			_, _ = fmt.Fprint(w, ": heartbeat\n\n")
			flusher.Flush()
		}
	}
}

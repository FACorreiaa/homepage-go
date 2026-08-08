package handler

import (
	"net/http"

	"myapp/internal/model"
	"myapp/internal/service"
	"myapp/ui/pages"
)

// homeFeaturedCount is how many case studies the landing page carries before
// handing off to /projects. Three is enough to establish range without
// turning the front door into the index.
const homeFeaturedCount = 3

type HomeHandler struct {
	Tracker *service.VisitTracker
}

func (h *HomeHandler) Show(w http.ResponseWriter, r *http.Request) {
	if err := pages.Home(h.Tracker.Snapshot(r.Context()), homeFeatured()).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

// homeFeatured reuses the single project list defined in projects.go rather
// than keeping a second copy in sync.
func homeFeatured() []model.ProjectItem {
	var out []model.ProjectItem
	for _, p := range allProjects {
		if !p.Featured {
			continue
		}
		out = append(out, p)
		if len(out) == homeFeaturedCount {
			break
		}
	}
	return out
}

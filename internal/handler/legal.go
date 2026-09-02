package handler

import (
	"net/http"

	"myapp/internal/model"
	"myapp/ui/pages"
)

func legalAppBySlug(slug string) (model.LegalApp, bool) {
	for _, a := range model.LegalApps() {
		if a.Slug == slug {
			return a, true
		}
	}
	return model.LegalApp{}, false
}

// LegalTerms serves GET /terms/{app}.
func LegalTerms(w http.ResponseWriter, r *http.Request) {
	app, ok := legalAppBySlug(r.PathValue("app"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	renderLegal(w, r, app, app.Terms)
}

// LegalPrivacy serves GET /privacy/{app}.
func LegalPrivacy(w http.ResponseWriter, r *http.Request) {
	app, ok := legalAppBySlug(r.PathValue("app"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	renderLegal(w, r, app, app.Privacy)
}

// LegalSupport serves GET /support/{app}: the page App Store Connect's
// "Support URL" points at.
func LegalSupport(w http.ResponseWriter, r *http.Request) {
	app, ok := legalAppBySlug(r.PathValue("app"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Cache-Control", "public, max-age=3600")
	if err := pages.LegalSupport(app).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func renderLegal(w http.ResponseWriter, r *http.Request, app model.LegalApp, doc model.LegalDocument) {
	w.Header().Set("Cache-Control", "public, max-age=3600")
	if err := pages.LegalPage(app, doc).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"myapp/internal/handler"

	"github.com/stretchr/testify/assert"
)

const (
	norviqWebsiteURL  = "https://norviq.org"
	norviqAppStoreURL = "https://apps.apple.com/pt/app/norviq/id6765849578?l=en-GB"
	khepriWebsiteURL  = "https://kheprios.com"
	lociWebsiteURL    = "https://lociai.fyi"
)

func TestProjectsList(t *testing.T) {
	r := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()
	handler.ProjectsList(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `href="`+norviqWebsiteURL+`"`)
	assert.Contains(t, body, `href="`+norviqAppStoreURL+`"`)
	assert.Contains(t, body, "norviq.org")
	assert.Contains(t, body, "Featured projects")
	assert.Contains(t, body, "Hosted products and systems I shipped and still stand behind.")
	// Uniform grid: no flagship hero, every featured card has the same anatomy.
	assert.NotContains(t, body, "project-feature-card--flagship")
	assert.Equal(t, 6, strings.Count(body, `class="project-feature-card project-showcase-card`),
		"hosted featured list is Norviq, Khepri, Loci, LuminaVault, HermesVault, Fandemic")
	assert.Contains(t, body, "Highlights")
	assert.Contains(t, body, "Holdings, watchlists, allocation with live pricing")
	assert.Contains(t, body, "Private financial command center, web + iPhone")
	assert.Contains(t, body, `href="/projects/luminavault"`)
	assert.Contains(t, body, `href="/projects/fandemic"`)
	assert.Contains(t, body, `href="/projects/khepri"`)
	assert.Contains(t, body, `href="/projects/loci"`)
	assert.Contains(t, body, `href="`+khepriWebsiteURL+`"`)
	assert.Contains(t, body, `href="`+lociWebsiteURL+`"`)
	assert.Contains(t, body, "kheprios.com")
	assert.Contains(t, body, "lociai.fyi")
}

func TestProjectDetail(t *testing.T) {
	tests := []struct {
		name string
		slug string
		want int
	}{
		{"known slug", "norviq", http.StatusOK},
		{"luminavault branded", "luminavault", http.StatusOK},
		{"khepri case study", "khepri", http.StatusOK},
		{"loci case study", "loci", http.StatusOK},
		{"unknown slug", "does-not-exist", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/projects/"+tt.slug, nil)
			r.SetPathValue("slug", tt.slug)
			w := httptest.NewRecorder()
			handler.ProjectDetail(w, r)
			assert.Equal(t, tt.want, w.Code)
			if tt.slug == "norviq" {
				assert.Contains(t, w.Body.String(), `href="`+norviqWebsiteURL+`"`)
				assert.Contains(t, w.Body.String(), `href="`+norviqAppStoreURL+`"`)
			}
			if tt.slug == "khepri" {
				body := w.Body.String()
				assert.Contains(t, body, `href="`+khepriWebsiteURL+`"`)
				assert.Contains(t, body, "/assets/static/projects/khepri/landing.webp")
				assert.Contains(t, body, "one coach, one memory")
			}
			if tt.slug == "loci" {
				body := w.Body.String()
				assert.Contains(t, body, `href="`+lociWebsiteURL+`"`)
				assert.Contains(t, body, "/assets/static/projects/loci/landing.webp")
				assert.Contains(t, body, "trip you can keep")
			}
		})
	}
}

func TestNorviqProductPage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /projects/{slug}", handler.ProjectDetail)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/projects/norviq", nil))
	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, "Nothing slips past.")
	assert.Contains(t, body, `id="screenshots"`)
	assert.Contains(t, body, "/assets/static/projects/norviq/home.webp")
	assert.Contains(t, body, `href="/privacy/norviq"`)
	assert.Contains(t, body, `href="/terms/norviq"`)
	assert.Contains(t, body, `href="/support/norviq"`)
	assert.Contains(t, body, `href="`+norviqAppStoreURL+`"`)
}

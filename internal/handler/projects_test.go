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
	assert.Contains(t, body, "project-feature-card--flagship")
	assert.Equal(t, 4, strings.Count(body, `class="project-feature-card project-showcase-card`),
		"hosted featured list stays Norviq, LuminaVault, HermesVault, Fandemic")
	assert.Contains(t, body, `href="/projects/luminavault"`)
	assert.Contains(t, body, `href="/projects/fandemic"`)
}

func TestProjectDetail(t *testing.T) {
	tests := []struct {
		name string
		slug string
		want int
	}{
		{"known slug", "norviq", http.StatusOK},
		{"luminavault branded", "luminavault", http.StatusOK},
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
		})
	}
}

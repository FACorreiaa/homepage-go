package handler_test

import (
	"net/http"
	"net/http/httptest"
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
	assert.Contains(t, w.Body.String(), `href="`+norviqWebsiteURL+`"`)
	assert.Contains(t, w.Body.String(), `href="`+norviqAppStoreURL+`"`)
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

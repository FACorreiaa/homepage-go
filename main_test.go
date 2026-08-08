package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The three revalidated assets are the ones that change without changing name:
// output.css is rebuilt on every deploy, and sw.js and manifest.json are the
// only levers for invalidating what the service worker already holds. Serving
// them immutable pins returning visitors to a year-old copy of the site.
func TestAssetCacheHeaders(t *testing.T) {
	const (
		revalidate = "public, max-age=0, must-revalidate"
		immutable  = "public, max-age=31536000, immutable"
	)

	tests := []struct {
		path string
		want string
	}{
		{"/assets/css/output.css", revalidate},
		{"/assets/static/sw.js", revalidate},
		{"/assets/static/manifest.json", revalidate},
		{"/assets/static/globe.js", immutable},
		{"/assets/fonts/geist/geist-variable.woff2", immutable},
		{"/assets/static/vendor/three/three.module.min.js", immutable},
	}

	t.Setenv("GO_ENV", "production")
	mux := http.NewServeMux()
	setupAssets(mux)

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			assert.Equal(t, tt.want, w.Header().Get("Cache-Control"))
		})
	}
}

func TestAssetCacheHeadersInDev(t *testing.T) {
	t.Setenv("GO_ENV", "development")
	mux := http.NewServeMux()
	setupAssets(mux)

	r := httptest.NewRequest("GET", "/assets/css/output.css", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	assert.Equal(t, "no-store", w.Header().Get("Cache-Control"))
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	revalidate = "public, max-age=0, must-revalidate"
	immutable  = "public, max-age=31536000, immutable"
)

// The three revalidated assets are the ones that change without changing name:
// output.css is rebuilt on every deploy, and sw.js and manifest.json are the
// only levers for invalidating what the service worker already holds. Serving
// them immutable pins returning visitors to a year-old copy of the site.
//
// This is a pure-function test on purpose. Asserting the header off a real
// response ties the test to which files happen to exist, and CI never builds
// output.css — net/http deletes Cache-Control when it serves a 404, so the
// assertion failed there for a reason that had nothing to do with the policy.
func TestAssetCacheControl(t *testing.T) {
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
		// The stripped path must never match: that was the original bug.
		{"css/output.css", immutable},
		{"static/sw.js", immutable},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equal(t, tt.want, assetCacheControl(tt.path, false))
			assert.Equal(t, "no-store", assetCacheControl(tt.path, true), "dev never caches")
		})
	}
}

// One end-to-end check that the policy is actually wired to the mux, on a file
// that is committed to the repo and so exists in every environment.
func TestAssetHandlerAppliesPolicy(t *testing.T) {
	t.Setenv("GO_ENV", "production")
	mux := http.NewServeMux()
	setupAssets(mux)

	r := httptest.NewRequest("GET", "/assets/static/sw.js", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, revalidate, w.Header().Get("Cache-Control"))
}

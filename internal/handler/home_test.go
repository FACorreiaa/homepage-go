package handler_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"myapp/assets"
	"myapp/internal/handler"

	"github.com/stretchr/testify/assert"
)

// The landing page is about to be rebuilt as a 3D scroll world. Everything the
// page exists to do lives in these assertions: the two forks a visitor can take,
// the featured work, and the telemetry that backs the claim above it. The scene
// is decoration; if any of this stops rendering, the rewrite broke the page.
//
// A nil tracker is a legal zero-traffic tracker — Snapshot guards its receiver —
// so the page renders without a database.
func renderHome(t *testing.T) string {
	t.Helper()
	h := &handler.HomeHandler{}
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.Show(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func TestHomeConversionPaths(t *testing.T) {
	body := renderHome(t)

	// The two forks, and the index they both fall back to.
	assert.Contains(t, body, `href="/proposal"`)
	assert.Contains(t, body, `href="/curriculum"`)
	assert.Contains(t, body, `href="/projects"`)

	assert.Contains(t, body, "YOU NEED A CONTRACTOR")
	assert.Contains(t, body, "YOU'RE HIRING")
}

func TestHomeFeaturedWork(t *testing.T) {
	body := renderHome(t)

	// Count card roots, not the --flagship modifier which also contains the class name.
	cards := strings.Count(body, `class="project-feature-card project-showcase-card`)
	assert.Equal(t, 3, cards, "landing page shows exactly three featured projects")
	assert.Equal(t, 3, strings.Count(body, "project-showcase-card"))

	// Flagship + 2-col grid: first card is full-bleed.
	assert.Contains(t, body, "project-feature-card--flagship")
	assert.Contains(t, body, "Featured projects")

	// Every card must link to its case study. Deep links are the whole point of
	// the section, and they are easy to lose when the markup is restructured.
	slugs := regexp.MustCompile(`href="/projects/([a-z0-9-]+)"`).FindAllStringSubmatch(body, -1)
	seen := map[string]bool{}
	for _, m := range slugs {
		seen[m[1]] = true
	}
	assert.GreaterOrEqual(t, len(seen), 3, "three distinct case-study links, got %v", seen)
}

func TestHomeOpsStrip(t *testing.T) {
	body := renderHome(t)

	// The strip is the evidence for the claim in the hero. Its labels are the
	// contract; the numbers behind them change every request.
	for _, label := range []string{"SERVING FROM", "UPTIME", "VIEWS 24H", "VISITORS NOW"} {
		assert.Contains(t, body, label)
	}
	assert.Contains(t, body, `data-live="true"`)
	assert.Contains(t, body, `href="/stats"`)
}

func TestHomeWorldScaffold(t *testing.T) {
	body := renderHome(t)

	// The canvas mounts into an empty, inert, aria-hidden container. If it is
	// ever given content or taken out of the fixed layer it starts affecting
	// the page, which is the one thing the scene must never do.
	assert.Contains(t, body, `<div class="world-stage" data-world aria-hidden="true"></div>`)
	assert.Contains(t, body, "world-shell")

	// Four stations, in the order the camera descends them.
	for _, station := range []string{"device", "server", "cluster", "horizon"} {
		assert.Contains(t, body, `data-world-station="`+station+`"`)
	}

	// gsap must load before ScrollTrigger, which registers itself against it.
	gsap := strings.Index(body, "vendor/gsap/gsap.min.js")
	scrollTrigger := strings.Index(body, "vendor/gsap/ScrollTrigger.min.js")
	assert.Positive(t, gsap)
	assert.Positive(t, scrollTrigger)
	assert.Less(t, gsap, scrollTrigger, "gsap must be ordered before ScrollTrigger")

	assert.Contains(t, body, "vendor/lenis/lenis.min.js")
	// The gopher module pulls in the GLB and GLTFLoader at runtime, so the page
	// only has to carry the module itself.

	// Fingerprinted, so match the path and assert the cache buster separately.
	assert.Regexp(t, `<script type="module" src="/assets/static/home-world\.js\?v=[0-9a-f]{10}">`, body)

	// Lenis owns scrolling here, so the two native behaviours that fight it are
	// stood down: scroll-smooth on <html>, and the scroll-top button.
	assert.NotContains(t, body, `<html lang="en" class="scroll-smooth">`)
	assert.Contains(t, body, "window.__lenis")
}

// The stylesheet URL must carry a content fingerprint. A browser that cached
// output.css while it was being served immutable will never revalidate the bare
// URL — no header can reach it, only a URL it has not seen. Shipping markup that
// depends on new CSS behind an unfingerprinted URL is what broke the live site.
func TestStylesheetIsFingerprinted(t *testing.T) {
	body := renderHome(t)
	assert.Regexp(t, `<link href="/assets/css/output\.css\?v=[0-9a-f]+" rel="stylesheet">`, body)
	assert.NotContains(t, body, `href="/assets/css/output.css"`,
		"the bare URL is unreachable in poisoned caches and must never be emitted")
}

func TestHomeRendersWithoutScripts(t *testing.T) {
	body := renderHome(t)

	// The 3D scene is progressive enhancement. Strip every script tag and the
	// page must still carry its content, exactly as it does for a crawler.
	withoutScripts := regexp.MustCompile(`(?s)<script.*?</script>`).ReplaceAllString(body, "")
	assert.Contains(t, withoutScripts, "One engineer, the whole stack, in production.")
	assert.Contains(t, withoutScripts, `href="/proposal"`)
	assert.Contains(t, withoutScripts, "project-showcase-card")
}

// The gopher mesh and its loader are the largest things the landing page pulls
// in, and they are exactly the kind of asset that gets pinned by accident. Only
// vendored libraries and fonts may be immutable; the model is neither.
func TestGopherAssetsAreNotPinned(t *testing.T) {
	for _, path := range []string{
		"/assets/static/models/gopher.glb",
		"/assets/static/gopher.js",
	} {
		assert.False(t, assets.IsImmutable(path), "%s must revalidate", path)
	}
	// The loader is version-pinned in its path, so it may be cached forever.
	assert.True(t, assets.IsImmutable("/assets/static/vendor/three/addons/loaders/GLTFLoader.js"))
}

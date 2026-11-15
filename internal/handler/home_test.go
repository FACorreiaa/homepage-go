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

// The cover is the first thing rendered and the only thing on the first screen,
// so it has to survive without the scene, the fonts, or anything else landing.
func TestHomeWordmarkCover(t *testing.T) {
	body := renderHome(t)

	assert.Contains(t, body, `<h1 class="wordmark">`)

	// The split is the whole idea: one name, two colours.
	assert.Contains(t, body, `<span class="wordmark-accent">Correia</span>`)

	// The wordmark comes before the hero line it introduces.
	assert.Less(t,
		strings.Index(body, `<h1 class="wordmark">`),
		strings.Index(body, `<h2 class="hero-title`),
		"wordmark must precede the hero line",
	)

	// Demoted, not deleted: the old hero line is now the sub-head.
	assert.Contains(t, body, `<h2 class="hero-title`)
	assert.Equal(t, 1, strings.Count(body, "<h1"), "exactly one h1 on the page")
}

// The ticker loops by travelling -50% across a track that holds its items
// twice. If the duplication ever stops happening the animation still runs, but
// it visibly snaps — so the doubling is the thing worth asserting.
func TestHomeMarqueeTrackIsDoubled(t *testing.T) {
	body := renderHome(t)

	assert.Contains(t, body, `class="marquee-track"`)
	assert.Equal(t, 2, strings.Count(body, ">GO</span>"), "marquee items must be rendered twice")
	assert.Equal(t, 2, strings.Count(body, ">KUBERNETES</span>"), "marquee items must be rendered twice")

	// Decoration, and it repeats itself — it must not be read out.
	assert.Contains(t, body, `<div class="marquee" aria-hidden="true">`)
}

// The fork panels became an index: the two conversion paths kept their copy and
// gained two navigation siblings.
func TestHomeIndexGrid(t *testing.T) {
	body := renderHome(t)

	assert.Equal(t, 4, strings.Count(body, `class="fork-panel fork-panel--`), "four index panels")

	// Each panel carries its own rule colour; a repeated modifier means two
	// panels are wearing the same one.
	for _, mod := range []string{"trail", "go", "signal", "ink"} {
		assert.Equal(t, 1, strings.Count(body, "fork-panel--"+mod), "one panel per rule colour: "+mod)
	}

	assert.Contains(t, body, `href="/blog"`)
	assert.Contains(t, body, `href="/bookmarks"`)
	assert.Contains(t, body, "WRITING")
	assert.Contains(t, body, "THE ARCHIVE")
}

func TestHomeFeaturedWork(t *testing.T) {
	body := renderHome(t)

	// Count card roots, not the --flagship modifier which also contains the class name.
	cards := strings.Count(body, `class="project-feature-card project-showcase-card`)
	assert.Equal(t, 4, cards, "landing page shows exactly four featured projects")
	assert.Equal(t, 4, strings.Count(body, "project-showcase-card"))

	// Uniform 2-col grid: the flagship hero is gone, so no card can outgrow its row.
	assert.NotContains(t, body, "project-feature-card--flagship")
	assert.Contains(t, body, "Featured projects")

	// Every card must link to its case study. Deep links are the whole point of
	// the section, and they are easy to lose when the markup is restructured.
	slugs := regexp.MustCompile(`href="/projects/([a-z0-9-]+)"`).FindAllStringSubmatch(body, -1)
	seen := map[string]bool{}
	for _, m := range slugs {
		seen[m[1]] = true
	}
	assert.GreaterOrEqual(t, len(seen), 4, "four distinct case-study links, got %v", seen)
	assert.True(t, seen["norviq"], "home featured includes Norviq, got %v", seen)
	assert.True(t, seen["khepri"], "home featured includes Khepri, got %v", seen)
	assert.True(t, seen["loci"], "home featured includes Loci, got %v", seen)
	assert.True(t, seen["luminavault"], "home featured includes LuminaVault, got %v", seen)
	// The live product link rides on the card, not only on the detail page.
	assert.Contains(t, body, `href="https://app.luminavault.fyi"`)
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

// The home page used to fly a three.js camera through the sections, with Lenis
// smooth scroll and GSAP ScrollTrigger driving it. It made the page feel slower
// than /projects for no gain, so it was removed. Keep it out: native scrolling,
// no WebGL, no scroll-hijacking libraries on the landing page.
func TestHomeHasNoScrollScene(t *testing.T) {
	body := renderHome(t)

	for _, gone := range []string{
		"world-stage", "data-world", "home-world.js",
		"vendor/gsap/", "vendor/lenis/", "window.__lenis",
		`<script type="importmap">`,
	} {
		assert.NotContains(t, body, gone)
	}
	assert.Contains(t, body, `<html lang="en" class="scroll-smooth">`)
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

	// Strip every script tag and the page must still carry its content,
	// exactly as it does for a crawler.
	withoutScripts := regexp.MustCompile(`(?s)<script.*?</script>`).ReplaceAllString(body, "")
	assert.Contains(t, withoutScripts, "One engineer, the whole stack, in production.")
	assert.Contains(t, withoutScripts, `href="/proposal"`)
	assert.Contains(t, withoutScripts, "project-showcase-card")
}

// The reel plays for everyone, so it must cost nothing until it is near the
// viewport: no autoplay attribute, no preload, and a cut per device class so a
// phone never downloads the 16:9 file.
func TestHomeShowreelIsLazyAndPerDevice(t *testing.T) {
	body := renderHome(t)

	assert.Contains(t, body, `data-showreel`)
	assert.Contains(t, body, `preload="none"`)
	assert.Contains(t, body, ` muted `)
	assert.Contains(t, body, `playsinline`)
	assert.NotRegexp(t, regexp.MustCompile(`<video[^>]*\sautoplay`), body)

	assert.Contains(t, body, `/assets/static/promo/reel-mobile.webm?v=`)
	assert.Contains(t, body, `/assets/static/promo/reel-desktop.mp4?v=`)
	assert.Contains(t, body, `media="(max-width: 767px)"`)

	// Every file the markup names must actually be embedded.
	for _, name := range []string{
		"static/promo/reel-desktop.webm", "static/promo/reel-desktop.mp4",
		"static/promo/reel-mobile.webm", "static/promo/reel-mobile.mp4",
		"static/promo/poster-desktop.webp", "static/promo/poster-mobile.webp",
		"static/showreel.js",
	} {
		_, err := assets.Assets.ReadFile(name)
		assert.NoError(t, err, name)
	}
}

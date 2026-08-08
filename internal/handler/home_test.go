package handler_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

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

	cards := strings.Count(body, "project-showcase-card")
	assert.Equal(t, 3, cards, "landing page shows exactly three featured projects")

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

func TestHomeRendersWithoutScripts(t *testing.T) {
	body := renderHome(t)

	// The 3D scene is progressive enhancement. Strip every script tag and the
	// page must still carry its content, exactly as it does for a crawler.
	withoutScripts := regexp.MustCompile(`(?s)<script.*?</script>`).ReplaceAllString(body, "")
	assert.Contains(t, withoutScripts, "One engineer, the whole stack, in production.")
	assert.Contains(t, withoutScripts, `href="/proposal"`)
	assert.Contains(t, withoutScripts, "project-showcase-card")
}

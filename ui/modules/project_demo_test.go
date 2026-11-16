package modules

import (
	"context"
	"strings"
	"testing"

	"myapp/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func renderDemo(t *testing.T, p model.ProjectItem) string {
	t.Helper()
	var sb strings.Builder
	require.NoError(t, ProjectDemo(p).Render(context.Background(), &sb))
	return sb.String()
}

func TestProjectDemoIsSilentWhenEmpty(t *testing.T) {
	body := renderDemo(t, model.ProjectItem{Title: "ScanItNow"})
	assert.Empty(t, strings.TrimSpace(body), "a project without a demo must render nothing")
}

func TestProjectDemoRendersVideo(t *testing.T) {
	body := renderDemo(t, model.ProjectItem{
		Title: "Norviq",
		Demo: model.ProjectDemo{
			Video:   "static/promo/norviq-demo.mp4",
			Poster:  "static/promo/norviq-demo-poster.webp",
			Caption: "A walkthrough.",
		},
	})

	assert.Contains(t, body, `src="/assets/static/promo/norviq-demo.mp4?v=`)
	assert.Contains(t, body, `poster="/assets/static/promo/norviq-demo-poster.webp?v=`)
	assert.Contains(t, body, "A walkthrough.")
	// A chosen video: user starts it, nothing downloads before that.
	assert.Contains(t, body, "controls")
	assert.Contains(t, body, `preload="none"`)
	assert.NotContains(t, body, "autoplay")
}

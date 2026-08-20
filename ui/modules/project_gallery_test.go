package modules

import (
	"context"
	"strings"
	"testing"

	"myapp/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func renderGallery(t *testing.T, p model.ProjectItem) string {
	t.Helper()
	var sb strings.Builder
	require.NoError(t, ProjectGallery(p).Render(context.Background(), &sb))
	return sb.String()
}

// Every project ships with an empty Gallery until real screenshots exist. The
// section must be absent in that state, not an empty frame or a bare heading —
// this is the same rule the other proof slots follow.
func TestProjectGalleryIsSilentWhenEmpty(t *testing.T) {
	body := renderGallery(t, model.ProjectItem{
		Title:        "Norviq",
		GalleryLabel: "App Store screenshots",
		// Gallery deliberately nil.
	})

	assert.Empty(t, strings.TrimSpace(body), "an empty gallery must render nothing at all")
}

func TestProjectGalleryRendersShots(t *testing.T) {
	body := renderGallery(t, model.ProjectItem{
		Title:        "Norviq",
		GalleryLabel: "App Store screenshots",
		Gallery: []model.GalleryShot{
			{Src: "/assets/static/projects/norviq-1.webp", Alt: "Holdings overview"},
			{Src: "/assets/static/projects/norviq-2.webp", Alt: "Research view"},
		},
		GalleryCaption: "Two moments from the iPhone app.",
	})

	assert.Contains(t, body, `src="/assets/static/projects/norviq-1.webp"`)
	assert.Contains(t, body, `alt="Research view"`)
	assert.Contains(t, body, "App Store screenshots")
	assert.Contains(t, body, "Two moments from the iPhone app.")

	// Below the fold on every project page, so nothing here races the hero.
	assert.Equal(t, 2, strings.Count(body, `loading="lazy"`))
}

// The label and caption are optional; supplying shots alone must still produce
// a valid section rather than an empty <p>.
func TestProjectGalleryOmitsBlankLabelAndCaption(t *testing.T) {
	body := renderGallery(t, model.ProjectItem{
		Title:   "Fandemic",
		Gallery: []model.GalleryShot{{Src: "/assets/static/projects/x.webp", Alt: "x"}},
	})

	assert.Contains(t, body, "Fandemic")
	assert.NotContains(t, body, "tracking-[0.16em]", "no label element when GalleryLabel is empty")
	assert.NotContains(t, body, "mt-8 max-w-xl", "no caption element when GalleryCaption is empty")
}

package modules

import (
	"context"
	"strings"
	"testing"

	"myapp/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func renderCard(t *testing.T, p model.ProjectItem) string {
	t.Helper()
	var sb strings.Builder
	require.NoError(t, ProjectFeatureCard(p, 0).Render(context.Background(), &sb))
	return sb.String()
}

func renderSection(t *testing.T, projects []model.ProjectItem, cta bool) string {
	t.Helper()
	var sb strings.Builder
	require.NoError(t, FeaturedProjectsSection(projects, "Featured projects", "", cta).Render(context.Background(), &sb))
	return sb.String()
}

// A card with nothing to highlight must not render an empty kicker — the block
// is all-or-nothing, like the gallery.
func TestFeatureCardOmitsEmptyHighlights(t *testing.T) {
	body := renderCard(t, model.ProjectItem{Slug: "x", Title: "X", Description: "d"})
	assert.NotContains(t, body, "Highlights")
	assert.NotContains(t, body, "project-highlights")
}

func TestFeatureCardCapsHighlightsAtThree(t *testing.T) {
	body := renderCard(t, model.ProjectItem{
		Slug: "x", Title: "X", Description: "d",
		Highlights: []string{"one", "two", "three", "four"},
	})
	assert.Contains(t, body, "Highlights")
	assert.Equal(t, 3, strings.Count(body, "<li"), "card shows at most three highlights")
	assert.NotContains(t, body, "four")
}

// The whole point of the redesign: the logo is a 72px mark inline with the
// title, not a hero centred in an art stage.
func TestFeatureCardRendersCompactLogo(t *testing.T) {
	body := renderCard(t, model.ProjectItem{
		Slug: "x", Title: "X", Description: "d", Tagline: "short pitch",
		LogoAsset: "/assets/static/projects/x.webp", Status: "Live", RoleTag: "Independent",
		DisplayGroup: "Web products", Tags: []string{"Go", "SwiftUI", "Postgres", "Redis"},
	})
	assert.Contains(t, body, `width="72"`)
	assert.NotContains(t, body, "project-art-stage")
	assert.Contains(t, body, "short pitch")
	assert.Contains(t, body, "[Web products]")
	assert.Contains(t, body, `href="/projects/x"`)
	// Three stack chips, not four.
	assert.Equal(t, 3, strings.Count(body, "project-meta-chip"))
	assert.NotContains(t, body, ">Redis<")
}

func TestFeatureCardFallsBackToInitials(t *testing.T) {
	body := renderCard(t, model.ProjectItem{Slug: "x", Title: "X", Description: "d", Icon: "HB"})
	assert.Contains(t, body, ">HB<")
	assert.NotContains(t, body, "<img")
}

func TestFeaturedSectionAllProjectsLink(t *testing.T) {
	projects := []model.ProjectItem{{Slug: "a", Title: "A"}, {Slug: "b", Title: "B"}}
	assert.NotContains(t, renderSection(t, projects, false), `href="/projects"`)
	assert.Contains(t, renderSection(t, projects, true), `href="/projects"`)
}

func TestFeaturedSectionHasNoFlagship(t *testing.T) {
	projects := []model.ProjectItem{{Slug: "a", Title: "A"}, {Slug: "b", Title: "B"}, {Slug: "c", Title: "C"}}
	body := renderSection(t, projects, false)
	assert.NotContains(t, body, "flagship")
	assert.Equal(t, 3, strings.Count(body, `class="project-feature-card project-showcase-card`))
}

package modules

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var navDestinations = []string{
	"/projects", "/about", "/curriculum", "/stack",
	"/blog", "/bookmarks", "/play", "/stats",
}

func renderNavbar(t *testing.T, active string) string {
	t.Helper()
	var sb strings.Builder
	require.NoError(t, Navbar(active).Render(context.Background(), &sb))
	return sb.String()
}

// Every destination must appear twice: once in the desktop bar and once in the
// drawer. A link added to one and forgotten in the other is invisible to half
// the visitors, which is how /stats shipped unreachable.
func TestNavbarHasEveryDestinationInBothVariants(t *testing.T) {
	body := renderNavbar(t, "home")
	for _, href := range navDestinations {
		assert.Equal(t, 2, strings.Count(body, `href="`+href+`"`),
			"%s should appear in both the desktop bar and the drawer", href)
	}
}

func TestNavbarMarksActivePage(t *testing.T) {
	// ActivePage values are set by the pages; an unconsumed one is dead wiring.
	for _, tc := range []struct{ active, href string }{
		{"projects", "/projects"},
		{"stats", "/stats"},
		{"bookmarks", "/bookmarks"},
	} {
		t.Run(tc.active, func(t *testing.T) {
			body := renderNavbar(t, tc.active)
			assert.Contains(t, body, "bg-accent", "active link should be styled")
		})
	}
}

// The breakpoint has regressed twice. The bar needs ~841px of its own width in
// the trimmed layout, so it cannot come back at md (768px), and it must not sit
// at lg (1024px) either — that collapses a 1000px-wide laptop window to a
// hamburger with room to spare.
func TestNavbarBreakpointIsMeasuredNotGuessed(t *testing.T) {
	body := renderNavbar(t, "home")

	assert.Contains(t, body, "min-[900px]:flex", "desktop bar reveals at the measured width")
	assert.Contains(t, body, "min-[900px]:hidden", "drawer hides at the same width")
	assert.NotContains(t, body, "lg:flex")
	assert.NotContains(t, body, "md:flex")

	// Below xl the bar buys room by dropping the verb, not the button, and by
	// dropping GitHub, which the footer carries on every page.
	assert.Contains(t, body, `<span class="hidden xl:inline">Request proposal</span>`)
	assert.Contains(t, body, `<span class="xl:hidden">Proposal</span>`)
	assert.Contains(t, body, "xl:inline-flex", "GitHub returns at xl")
}

package pages

import (
	"strings"
	"testing"

	"myapp/internal/model"
)

func TestNorviqExtraJSONLDIncludesFAQ(t *testing.T) {
	raw := norviqExtraJSONLD(model.ProjectItem{Slug: "norviq", Title: "Norviq", Description: "d", LogoAsset: "/assets/static/projects/norviq-icon.webp"})
	if strings.Count(raw, `<script type="application/ld+json">`) != 2 {
		t.Fatalf("want two script tags, got %q", raw)
	}
	if !strings.Contains(raw, `"@type":"FAQPage"`) {
		t.Fatalf("missing FAQPage: %s", raw)
	}
	if !strings.Contains(raw, "Is Norviq free?") {
		t.Fatalf("missing FAQ question: %s", raw)
	}
}

func TestNorviqFAQsStayInSyncWithJSONLD(t *testing.T) {
	faqs := norviqFAQs()
	if len(faqs) != 6 {
		t.Fatalf("len=%d", len(faqs))
	}
}

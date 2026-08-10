package layouts

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestParseBlogLastMod(t *testing.T) {
	tests := []struct {
		updated, date, want string
	}{
		{"2026-07-01 12:00:00", "2024-01-15", "2026-07-01"},
		{"", "2024-01-15", "2024-01-15"},
		{"", "January 2, 2025", "2025-01-02"},
		{"", "", ""},
	}
	for _, tt := range tests {
		got := ParseBlogLastMod(tt.updated, tt.date)
		if got != tt.want {
			t.Errorf("ParseBlogLastMod(%q,%q)=%q want %q", tt.updated, tt.date, got, tt.want)
		}
	}
}

func TestBlogPostJSONLD(t *testing.T) {
	raw := BlogPostJSONLD(
		"norviq",
		`Hello "world" & <friends>`,
		"A summary",
		"Engineering",
		"2025-03-01",
		"2025-04-01 10:00:00",
		"",
	)
	if !strings.HasPrefix(raw, `<script type="application/ld+json">`) || !strings.HasSuffix(raw, `</script>`) {
		t.Fatalf("not a script tag: %s", raw)
	}
	inner := strings.TrimSuffix(strings.TrimPrefix(raw, `<script type="application/ld+json">`), `</script>`)
	var payload map[string]any
	if err := json.Unmarshal([]byte(inner), &payload); err != nil {
		t.Fatal(err)
	}
	graph, _ := payload["@graph"].([]any)
	if len(graph) != 2 {
		t.Fatalf("graph len=%d", len(graph))
	}
	post, _ := graph[0].(map[string]any)
	if post["@type"] != "BlogPosting" {
		t.Fatalf("type=%v", post["@type"])
	}
	if post["headline"] != `Hello "world" & <friends>` {
		t.Fatalf("headline escaped wrong: %v", post["headline"])
	}
	if post["datePublished"] != "2025-03-01" {
		t.Fatalf("datePublished=%v", post["datePublished"])
	}
	if post["dateModified"] != "2025-04-01" {
		t.Fatalf("dateModified=%v", post["dateModified"])
	}
	crumbs, _ := graph[1].(map[string]any)
	if crumbs["@type"] != "BreadcrumbList" {
		t.Fatalf("crumbs type=%v", crumbs["@type"])
	}
}

func TestProjectJSONLD(t *testing.T) {
	raw := ProjectJSONLD("luminavault", "LuminaVault", "Second brain", "/assets/static/projects/luminavault-icon.webp")
	inner := strings.TrimSuffix(strings.TrimPrefix(raw, `<script type="application/ld+json">`), `</script>`)
	var payload map[string]any
	if err := json.Unmarshal([]byte(inner), &payload); err != nil {
		t.Fatal(err)
	}
	graph, _ := payload["@graph"].([]any)
	work, _ := graph[0].(map[string]any)
	if work["@type"] != "CreativeWork" {
		t.Fatalf("type=%v", work["@type"])
	}
	if work["image"] != "https://facorreia.com/assets/static/projects/luminavault-icon.webp" {
		t.Fatalf("image=%v", work["image"])
	}
}

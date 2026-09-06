package layouts

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAbsoluteAssetURL(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, in, want string
	}{
		{"default", "", "https://facorreia.com/assets/static/og.png"},
		{"relative", "/assets/static/projects/norviq/home.webp", "https://facorreia.com/assets/static/projects/norviq/home.webp"},
		{"absolute", "https://cdn.example/x.png", "https://cdn.example/x.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := absoluteAssetURL(tt.in); got != tt.want {
				t.Fatalf("got %q want %q", got, tt.want)
			}
		})
	}
}

func TestOgImageDefaultAndRobots(t *testing.T) {
	t.Parallel()
	if got := (LayoutProps{}).ogImage(); got != "https://facorreia.com/assets/static/og.png" {
		t.Fatalf("default ogImage=%q", got)
	}
	if got := (LayoutProps{OGImage: "/assets/x.png"}).ogImage(); got != "https://facorreia.com/assets/x.png" {
		t.Fatalf("relative ogImage=%q", got)
	}
	if got := (LayoutProps{}).robots(); got != "index,follow,max-image-preview:large" {
		t.Fatalf("robots=%q", got)
	}
	if got := (LayoutProps{NoIndex: true}).robots(); got != "noindex,nofollow" {
		t.Fatalf("noindex robots=%q", got)
	}
}

func TestSiteGraphJSONLDIncludesLinkedInAndShareCard(t *testing.T) {
	t.Parallel()
	raw := SiteGraphJSONLD()
	inner := strings.TrimSuffix(strings.TrimPrefix(raw, `<script type="application/ld+json">`), `</script>`)
	var payload map[string]any
	if err := json.Unmarshal([]byte(inner), &payload); err != nil {
		t.Fatal(err)
	}
	graph, _ := payload["@graph"].([]any)
	if len(graph) != 3 {
		t.Fatalf("graph len=%d", len(graph))
	}
	person, _ := graph[1].(map[string]any)
	sameAs, _ := person["sameAs"].([]any)
	joined := ""
	for _, v := range sameAs {
		joined += v.(string) + " "
	}
	if !strings.Contains(joined, "linkedin.com/in/fernando-correia") || !strings.Contains(joined, "github.com/FACorreiaa") {
		t.Fatalf("sameAs=%v", sameAs)
	}
	if person["image"] != "https://facorreia.com/assets/static/og.png" {
		t.Fatalf("person image=%v", person["image"])
	}
}

func TestFAQPageJSONLD(t *testing.T) {
	t.Parallel()
	raw := FAQPageJSONLD("/projects/norviq", []FAQItem{
		{Question: "Is Norviq free?", Answer: "Yes."},
	})
	inner := strings.TrimSuffix(strings.TrimPrefix(raw, `<script type="application/ld+json">`), `</script>`)
	var payload map[string]any
	if err := json.Unmarshal([]byte(inner), &payload); err != nil {
		t.Fatal(err)
	}
	if payload["@type"] != "FAQPage" {
		t.Fatalf("type=%v", payload["@type"])
	}
	main, _ := payload["mainEntity"].([]any)
	if len(main) != 1 {
		t.Fatalf("mainEntity len=%d", len(main))
	}
	q, _ := main[0].(map[string]any)
	if q["name"] != "Is Norviq free?" {
		t.Fatalf("question=%v", q["name"])
	}
}

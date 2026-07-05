package service

import (
	"testing"
	"testing/fstest"
	"time"
)

func testVault() fstest.MapFS {
	return fstest.MapFS{
		"raw/Dev/go-generics.md": &fstest.MapFile{
			Data:    []byte("---\ntitle: Go Generics Deep Dive\nurl: https://go.dev/blog/generics\ndate: 2026-03-10\n---\nGenerics change how Go libraries are designed. See [[Rust Traits]] for contrast."),
			ModTime: time.Date(2026, 3, 11, 0, 0, 0, 0, time.UTC),
		},
		"raw/Dev/rust-traits.md": &fstest.MapFile{
			Data:    []byte("---\ningested_at: 2026-02-01T10:00:00Z\n---\nTraits are Rust's interfaces."),
			ModTime: time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC),
		},
		"raw/AI/2026-04-05-llm-agents.md": &fstest.MapFile{
			Data:    []byte("Agents plan and act in loops."),
			ModTime: time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC),
		},
		"raw/AI/no-date-note.md": &fstest.MapFile{
			Data:    []byte("Fallback to mtime."),
			ModTime: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		},
		"raw/AI/skip.sync-conflict-1.md": &fstest.MapFile{Data: []byte("ignored")},
		"wiki/ai-scoreboard.md": &fstest.MapFile{
			Data:    []byte("Weekly model rankings."),
			ModTime: time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		"daily-note.md": &fstest.MapFile{
			Data:    []byte("Root-level note."),
			ModTime: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"Diary/private-entry.md": &fstest.MapFile{
			Data: []byte("never public"),
		},
		".obsidian/workspace.md": &fstest.MapFile{
			Data: []byte("editor state"),
		},
		"_AGENT_INDEX.md": &fstest.MapFile{
			Data: []byte("Vault root: /mnt/volume-fsn1-1/hermes-data"),
		},
		"AI/INDEX.md": &fstest.MapFile{
			Data: []byte("generated index"),
		},
		"AI/README.md": &fstest.MapFile{
			Data: []byte("folder readme"),
		},
	}
}

func TestBuildBookmarkIndexEnrichment(t *testing.T) {
	idx := BuildBookmarkIndex(testVault())

	if len(idx.Files) != 6 {
		t.Fatalf("want 6 files, got %d", len(idx.Files))
	}
	if idx.BySlug["private-entry"] != nil || idx.BySlug["workspace"] != nil {
		t.Fatal("private/hidden dirs must not be indexed")
	}
	if idx.BySlug["_agent_index"] != nil || idx.BySlug["index"] != nil || idx.BySlug["readme"] != nil {
		t.Fatal("generated scaffolding files must not be indexed")
	}
	if w := idx.BySlug["ai-scoreboard"]; w == nil || w.Category != "wiki" {
		t.Errorf("wiki category: %+v", w)
	}
	if r := idx.BySlug["daily-note"]; r == nil || r.Category != "Notes" {
		t.Errorf("root files should get Notes category: %+v", r)
	}

	gg := idx.BySlug["go-generics"]
	if gg == nil {
		t.Fatal("go-generics not indexed")
	}
	if gg.Name != "Go Generics Deep Dive" {
		t.Errorf("frontmatter title not used: %q", gg.Name)
	}
	if gg.SourceURL != "https://go.dev/blog/generics" {
		t.Errorf("source url: %q", gg.SourceURL)
	}
	if gg.SourceDomain() != "go.dev" {
		t.Errorf("source domain: %q", gg.SourceDomain())
	}
	if gg.Date.Format("2006-01-02") != "2026-03-10" {
		t.Errorf("frontmatter date not used: %v", gg.Date)
	}
	if len(gg.Links) != 1 || gg.Links[0] != "rust-traits" {
		t.Errorf("wikilinks: %v", gg.Links)
	}
	if gg.Excerpt == "" || gg.Excerpt[:8] != "Generics" {
		t.Errorf("excerpt: %q", gg.Excerpt)
	}

	rt := idx.BySlug["rust-traits"]
	if rt.Date.Format("2006-01-02") != "2026-02-01" {
		t.Errorf("ingested_at date not used: %v", rt.Date)
	}

	agents := idx.BySlug["2026-04-05-llm-agents"]
	if agents.Date.Format("2006-01-02") != "2026-04-05" {
		t.Errorf("filename date prefix not used: %v", agents.Date)
	}

	nd := idx.BySlug["no-date-note"]
	if nd.Date.Format("2006-01-02") != "2026-01-15" {
		t.Errorf("mtime fallback not used: %v", nd.Date)
	}

	// Date-desc ordering (ai-scoreboard has the newest mtime: 2026-05-01)
	if idx.Files[0].Slug != "ai-scoreboard" {
		t.Errorf("expected newest first, got %s", idx.Files[0].Slug)
	}
}

func TestFilter(t *testing.T) {
	idx := BuildBookmarkIndex(testVault())

	items, total := idx.Filter("generics", "", "", 1, 30)
	if total != 1 || items[0].Slug != "go-generics" {
		t.Errorf("search: total=%d items=%v", total, items)
	}

	_, total = idx.Filter("", "AI", "", 1, 30)
	if total != 2 {
		t.Errorf("category filter: total=%d", total)
	}

	items, total = idx.Filter("", "", "name", 1, 2)
	if total != 6 || len(items) != 2 {
		t.Errorf("pagination: total=%d len=%d", total, len(items))
	}
	if items[0].Name > items[1].Name {
		t.Errorf("name sort broken: %q > %q", items[0].Name, items[1].Name)
	}

	items, _ = idx.Filter("", "", "", 4, 2)
	if len(items) != 0 {
		t.Errorf("out-of-range page should be empty, got %d", len(items))
	}
}

func TestCategoryCountsAndRelated(t *testing.T) {
	idx := BuildBookmarkIndex(testVault())

	counts := idx.CategoryCounts()
	if len(counts) != 4 || counts[0].Name != "AI" || counts[0].Count != 2 {
		t.Errorf("category counts: %v", counts)
	}

	gg := idx.BySlug["go-generics"]
	related := idx.Related(gg, 3)
	if len(related) == 0 || related[0].Slug != "rust-traits" {
		t.Errorf("wikilink neighbor should rank first: %v", related)
	}
	for _, r := range related {
		if r.Slug == gg.Slug {
			t.Error("related contains self")
		}
	}
}

func TestMakeExcerpt(t *testing.T) {
	body := "# Heading\n\n![img](http://x.com/a.png)\n\nSome **bold** text with [a link](http://example.com) and [[Wiki Page|display]] plus <b>html</b>.\n\n```\ncode block\n```"
	got := makeExcerpt(body, 40)
	for _, banned := range []string{"#", "![", "](", "<b>", "```", "http"} {
		if contains := indexOf(got, banned); contains >= 0 {
			t.Errorf("excerpt contains %q: %q", banned, got)
		}
	}
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

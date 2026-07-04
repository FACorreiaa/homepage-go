package service

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
)

type VaultFile struct {
	Slug      string
	Name      string
	Category  string
	FullPath  string
	Excerpt   string
	SourceURL string
	Date      time.Time
	WordCount int
	Links     []string // wikilink target slugs
}

func (f VaultFile) SourceDomain() string {
	u := f.SourceURL
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")
	u = strings.TrimPrefix(u, "www.")
	if i := strings.IndexByte(u, '/'); i != -1 {
		u = u[:i]
	}
	return u
}

func (f VaultFile) ReadMinutes() int {
	m := f.WordCount / 220
	if m < 1 {
		m = 1
	}
	return m
}

type SidebarItem struct {
	Slug string
	Name string
}

type SidebarCategory struct {
	Name  string
	Items []SidebarItem
}

type CategoryCount struct {
	Name  string
	Count int
}

type GraphNode struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Val      int    `json:"val"`
	Group    int    `json:"group"`
	Category string `json:"category"`
	IsHub    bool   `json:"isHub"`
}

type GraphLink struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type GraphData struct {
	Nodes []GraphNode `json:"nodes"`
	Links []GraphLink `json:"links"`
}

type BookmarkIndex struct {
	Files     []VaultFile // sorted by Date desc
	BySlug    map[string]*VaultFile
	Sidebar   []SidebarCategory
	GraphData GraphData
	BuiltAt   time.Time
}

var wikilinkRe = regexp.MustCompile(`\[\[([^\]|]+)(?:\|[^\]]+)?\]\]`)

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

// vaultMeta covers the frontmatter fields that actually occur in the vault
// (source/url/date variants differ by ingestion pipeline).
type vaultMeta struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Source      string `yaml:"source"`
	SourceURL   string `yaml:"source_url"`
	URL         string `yaml:"url"`
	MalURL      string `yaml:"mal_url"`
	Date        string `yaml:"date"`
	Published   string `yaml:"published"`
	PubDate     string `yaml:"pubDate"`
	Created     string `yaml:"created"`
	IngestedAt  string `yaml:"ingested_at"`
	SyncedAt    string `yaml:"synced_at"`
}

func (m vaultMeta) sourceURL() string {
	for _, c := range []string{m.URL, m.SourceURL, m.Source, m.MalURL} {
		if strings.HasPrefix(c, "http://") || strings.HasPrefix(c, "https://") {
			return c
		}
	}
	return ""
}

// fallbackFrontmatter scans "key: value" lines of a frontmatter block that
// failed strict YAML parsing (unquoted colons in values are common in
// clipped notes).
func fallbackFrontmatter(content string) vaultMeta {
	var meta vaultMeta
	if !strings.HasPrefix(content, "---") {
		return meta
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "---" {
			break
		}
		key, val, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		val = strings.Trim(strings.TrimSpace(val), `"'`)
		switch strings.TrimSpace(key) {
		case "title":
			meta.Title = val
		case "description":
			meta.Description = val
		case "source":
			meta.Source = val
		case "source_url":
			meta.SourceURL = val
		case "url":
			meta.URL = val
		case "date":
			meta.Date = val
		case "published":
			meta.Published = val
		case "created":
			meta.Created = val
		case "ingested_at":
			meta.IngestedAt = val
		case "synced_at":
			meta.SyncedAt = val
		}
	}
	return meta
}

var dateLayouts = []string{
	time.RFC3339Nano, time.RFC3339,
	"2006-01-02 15:04:05", "2006-01-02T15:04:05", "2006-01-02",
	"Mon, 02 Jan 2006 15:04:05 -0700", "Mon, 02 Jan 2006 15:04:05 MST",
}

func parseAnyDate(s string) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	for _, l := range dateLayouts {
		if t, err := time.Parse(l, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

var filenameDateRe = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})[-_]`)

func (m vaultMeta) date(baseName string, mtime time.Time) time.Time {
	for _, c := range []string{m.Date, m.Published, m.PubDate, m.Created, m.IngestedAt, m.SyncedAt} {
		if t, ok := parseAnyDate(c); ok {
			return t
		}
	}
	if dm := filenameDateRe.FindStringSubmatch(baseName); dm != nil {
		if t, err := time.Parse("2006-01-02", dm[1]); err == nil {
			return t
		}
	}
	return mtime
}

var (
	codeBlockRe = regexp.MustCompile("(?s)```.*?```")
	htmlTagRe   = regexp.MustCompile(`<[^>]*>`)
	imageRe     = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)`)
	mdLinkRe    = regexp.MustCompile(`\[([^\]]*)\]\([^)]*\)`)
	bareURLRe   = regexp.MustCompile(`https?://\S+`)
	mdSyntaxRe  = regexp.MustCompile("[#*_`>|]+")
	spaceRe     = regexp.MustCompile(`\s+`)
)

func makeExcerpt(body string, maxWords int) string {
	s := codeBlockRe.ReplaceAllString(body, " ")
	s = imageRe.ReplaceAllString(s, " ")
	s = mdLinkRe.ReplaceAllString(s, "$1")
	s = wikilinkRe.ReplaceAllStringFunc(s, func(m string) string {
		sub := wikilinkRe.FindStringSubmatch(m)
		if len(sub) > 1 {
			return sub[1]
		}
		return " "
	})
	s = htmlTagRe.ReplaceAllString(s, " ")
	s = bareURLRe.ReplaceAllString(s, " ")
	s = mdSyntaxRe.ReplaceAllString(s, " ")
	s = spaceRe.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	if len(words) > maxWords {
		return strings.Join(words[:maxWords], " ") + "…"
	}
	return strings.Join(words, " ")
}

func BuildBookmarkIndex(vaultFS fs.FS) BookmarkIndex {
	var files []VaultFile
	seen := map[string]bool{}

	fs.WalkDir(vaultFS, "raw", func(path string, d fs.DirEntry, err error) error { //nolint:errcheck
		if err != nil || d.IsDir() {
			return nil
		}
		name := d.Name()
		if !strings.HasSuffix(name, ".md") {
			return nil
		}
		if strings.Contains(name, ".sync-conflict") || strings.HasPrefix(name, ".syncthing.") || strings.HasSuffix(name, ".tmp") {
			return nil
		}
		baseName := strings.TrimSuffix(name, ".md")
		slug := slugify(baseName)
		if seen[slug] {
			return nil
		}

		data, rerr := fs.ReadFile(vaultFS, path)
		if rerr != nil {
			return nil
		}
		seen[slug] = true

		rel := strings.TrimPrefix(path, "raw/")
		parts := strings.SplitN(rel, "/", 2)
		category := "Uncategorized"
		if len(parts) > 1 {
			category = parts[0]
		}

		var meta vaultMeta
		body, ferr := frontmatter.Parse(strings.NewReader(string(data)), &meta)
		var bodyStr string
		if ferr == nil {
			bodyStr = string(body)
		} else {
			// Clipped notes often have YAML that strict parsers reject
			// (e.g. unquoted colons in titles); salvage with a line scan.
			meta = fallbackFrontmatter(string(data))
			bodyStr = StripFrontmatter(string(data))
		}

		var mtime time.Time
		if info, serr := d.Info(); serr == nil {
			mtime = info.ModTime()
		}

		title := strings.TrimSpace(meta.Title)
		if title == "" {
			title = baseName
		}

		excerpt := strings.TrimSpace(meta.Description)
		if excerpt == "" {
			excerpt = makeExcerpt(bodyStr, 40)
		}

		var links []string
		for _, m := range wikilinkRe.FindAllStringSubmatch(bodyStr, -1) {
			links = append(links, slugify(m[1]))
		}

		files = append(files, VaultFile{
			Slug:      slug,
			Name:      title,
			Category:  category,
			FullPath:  path,
			Excerpt:   excerpt,
			SourceURL: meta.sourceURL(),
			Date:      meta.date(baseName, mtime),
			WordCount: len(strings.Fields(bodyStr)),
			Links:     links,
		})
		return nil
	})

	sort.Slice(files, func(i, j int) bool {
		if !files[i].Date.Equal(files[j].Date) {
			return files[i].Date.After(files[j].Date)
		}
		return files[i].Name < files[j].Name
	})

	bySlug := make(map[string]*VaultFile, len(files))
	for i := range files {
		bySlug[files[i].Slug] = &files[i]
	}

	// Sidebar (alphabetical categories and items)
	grouped := map[string][]SidebarItem{}
	for _, f := range files {
		grouped[f.Category] = append(grouped[f.Category], SidebarItem{Slug: f.Slug, Name: f.Name})
	}
	sidebar := make([]SidebarCategory, 0, len(grouped))
	for cat, items := range grouped {
		sort.Slice(items, func(i, j int) bool { return items[i].Name < items[j].Name })
		sidebar = append(sidebar, SidebarCategory{Name: cat, Items: items})
	}
	sort.Slice(sidebar, func(i, j int) bool { return sidebar[i].Name < sidebar[j].Name })

	// Graph
	categoryIndex := map[string]int{}
	catCount := map[string]int{}
	for _, f := range files {
		if _, ok := categoryIndex[f.Category]; !ok {
			categoryIndex[f.Category] = len(categoryIndex)
		}
		catCount[f.Category]++
	}

	nodeIDs := map[string]bool{}
	var nodes []GraphNode
	var links []GraphLink

	for cat, grp := range categoryIndex {
		hubID := "hub-" + slugify(cat)
		nodes = append(nodes, GraphNode{ID: hubID, Name: cat, Val: max(3, catCount[cat]/2), Group: grp, Category: cat, IsHub: true})
		nodeIDs[hubID] = true
	}
	for _, f := range files {
		grp := categoryIndex[f.Category]
		nodes = append(nodes, GraphNode{ID: f.Slug, Name: f.Name, Val: 1, Group: grp, Category: f.Category})
		nodeIDs[f.Slug] = true
		links = append(links, GraphLink{Source: f.Slug, Target: "hub-" + slugify(f.Category)})
	}
	for _, f := range files {
		for _, target := range f.Links {
			if nodeIDs[target] && target != f.Slug {
				links = append(links, GraphLink{Source: f.Slug, Target: target})
			}
		}
	}

	return BookmarkIndex{
		Files:     files,
		BySlug:    bySlug,
		Sidebar:   sidebar,
		GraphData: GraphData{Nodes: nodes, Links: links},
		BuiltAt:   time.Now(),
	}
}

// CrossLinkCount is the number of wikilink edges between notes (the graph
// also holds one structural edge per note to its category hub).
func (idx BookmarkIndex) CrossLinkCount() int {
	n := len(idx.GraphData.Links) - len(idx.Files)
	if n < 0 {
		n = 0
	}
	return n
}

// Filter returns one page of notes matching a search query and/or category,
// plus the total match count. sortKey: "date" (default), "name", "category".
func (idx BookmarkIndex) Filter(q, category, sortKey string, page, perPage int) ([]VaultFile, int) {
	q = strings.ToLower(strings.TrimSpace(q))
	var matched []VaultFile
	for _, f := range idx.Files {
		if category != "" && f.Category != category {
			continue
		}
		if q != "" &&
			!strings.Contains(strings.ToLower(f.Name), q) &&
			!strings.Contains(strings.ToLower(f.Excerpt), q) {
			continue
		}
		matched = append(matched, f)
	}

	switch sortKey {
	case "name":
		sort.Slice(matched, func(i, j int) bool { return matched[i].Name < matched[j].Name })
	case "category":
		sort.Slice(matched, func(i, j int) bool {
			if matched[i].Category != matched[j].Category {
				return matched[i].Category < matched[j].Category
			}
			return matched[i].Name < matched[j].Name
		})
	default:
		// idx.Files is already date-desc
	}

	total := len(matched)
	if perPage <= 0 {
		perPage = 30
	}
	if page < 1 {
		page = 1
	}
	start := (page - 1) * perPage
	if start >= total {
		return nil, total
	}
	end := start + perPage
	if end > total {
		end = total
	}
	return matched[start:end], total
}

// CategoryCounts returns categories sorted by note count (desc).
func (idx BookmarkIndex) CategoryCounts() []CategoryCount {
	counts := map[string]int{}
	for _, f := range idx.Files {
		counts[f.Category]++
	}
	out := make([]CategoryCount, 0, len(counts))
	for name, n := range counts {
		out = append(out, CategoryCount{Name: name, Count: n})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		return out[i].Name < out[j].Name
	})
	return out
}

// Related returns up to limit notes related to f: wikilink neighbors first,
// then recent notes from the same category.
func (idx BookmarkIndex) Related(f *VaultFile, limit int) []VaultFile {
	if f == nil || limit <= 0 {
		return nil
	}
	seen := map[string]bool{f.Slug: true}
	var out []VaultFile
	add := func(v *VaultFile) {
		if v != nil && !seen[v.Slug] && len(out) < limit {
			seen[v.Slug] = true
			out = append(out, *v)
		}
	}
	for _, slug := range f.Links {
		add(idx.BySlug[slug])
	}
	// Reverse links (notes pointing at this one)
	for i := range idx.Files {
		if len(out) >= limit {
			break
		}
		for _, l := range idx.Files[i].Links {
			if l == f.Slug {
				add(&idx.Files[i])
				break
			}
		}
	}
	for i := range idx.Files {
		if len(out) >= limit {
			break
		}
		if idx.Files[i].Category == f.Category {
			add(&idx.Files[i])
		}
	}
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ResolveWikilinks(content string) string {
	return wikilinkRe.ReplaceAllStringFunc(content, func(match string) string {
		inner := wikilinkRe.FindStringSubmatch(match)
		if len(inner) < 2 {
			return match
		}
		target := inner[1]
		slug := slugify(target)
		display := target
		if pipe := strings.Index(inner[0], "|"); pipe != -1 {
			display = inner[0][pipe+1 : len(inner[0])-2]
		}
		return `<a href="/bookmarks/` + slug + `" class="text-primary underline underline-offset-2">` + display + `</a>`
	})
}

func StripFrontmatter(content string) string {
	if !strings.HasPrefix(content, "---") {
		return content
	}
	lines := strings.Split(content, "\n")
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			return strings.Join(lines[i+1:], "\n")
		}
	}
	return content
}

func FileExt(path string) string {
	return filepath.Ext(path)
}

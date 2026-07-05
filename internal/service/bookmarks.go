package service

import (
	"hash/fnv"
	"io/fs"
	"math"
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
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Val      float64 `json:"val"`
	Group    int     `json:"group"`
	Category string  `json:"category"`
	IsHub    bool    `json:"isHub"`
	// Precomputed, deterministic 3D position (fx/fy/fz are treated as fixed
	// coordinates by 3d-force-graph, so no client-side physics is needed and
	// the layout never jitters between loads).
	FX float64 `json:"fx"`
	FY float64 `json:"fy"`
	FZ float64 `json:"fz"`
	// Heat is recency in [0,1]: 1 = touched today, decaying over months.
	Heat float64 `json:"heat"`
	Deg  int     `json:"-"`
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

var filenameDateRe = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

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

// privateDirs are vault directories that must never reach the public site,
// enforced here as well as in the rsync excludes (defense in depth).
var privateDirs = map[string]bool{
	"diary":   true,
	"kanban":  true,
	"tickets": true,
	"setup":   true,
	"scripts": true,
}

func BuildBookmarkIndex(vaultFS fs.FS) BookmarkIndex {
	var files []VaultFile
	seen := map[string]bool{}

	fs.WalkDir(vaultFS, ".", func(path string, d fs.DirEntry, err error) error { //nolint:errcheck
		if err != nil {
			return nil
		}
		name := d.Name()
		if d.IsDir() {
			if path == "." {
				return nil
			}
			if strings.HasPrefix(name, ".") || privateDirs[strings.ToLower(name)] {
				return fs.SkipDir
			}
			return nil
		}
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

		// Category: top-level directory; inside raw/ the second level is the
		// real category (raw/AI/..., raw/Dev/...). Root-level files → Notes.
		parts := strings.Split(path, "/")
		category := "Notes"
		switch {
		case len(parts) == 1:
			// root-level file
		case strings.EqualFold(parts[0], "raw"):
			if len(parts) > 2 {
				category = parts[1]
			} else {
				category = "Uncategorized"
			}
		default:
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

	// Graph: category hubs on a fibonacci sphere, notes clustered around
	// their hub with hash-deterministic offsets. Node size encodes note
	// length + connectivity; heat encodes recency.
	categoryIndex := map[string]int{}
	catCount := map[string]int{}
	for _, f := range files {
		if _, ok := categoryIndex[f.Category]; !ok {
			categoryIndex[f.Category] = len(categoryIndex)
		}
		catCount[f.Category]++
	}

	const hubRadius = 1500.0
	nCats := len(categoryIndex)
	hubPos := map[string][3]float64{}
	for cat, grp := range categoryIndex {
		x, y, z := fibonacciSphere(grp, nCats)
		hubPos[cat] = [3]float64{x * hubRadius, y * hubRadius, z * hubRadius}
	}

	// Degree counts only edges between notes that actually exist — clipped
	// notes often carry hundreds of [[entity]] links with no matching note.
	existing := map[string]bool{}
	for _, f := range files {
		existing[f.Slug] = true
	}
	degree := map[string]int{}
	for _, f := range files {
		for _, target := range f.Links {
			if existing[target] && target != f.Slug {
				degree[target]++
				degree[f.Slug]++
			}
		}
	}

	now := time.Now()
	nodeIDs := map[string]bool{}
	var nodes []GraphNode
	var links []GraphLink

	for cat, grp := range categoryIndex {
		hubID := "hub-" + slugify(cat)
		p := hubPos[cat]
		val := math.Max(6, math.Sqrt(float64(catCount[cat]))*1.6)
		nodes = append(nodes, GraphNode{
			ID: hubID, Name: cat, Val: val, Group: grp, Category: cat, IsHub: true,
			FX: p[0], FY: p[1], FZ: p[2], Heat: 0.55,
		})
		nodeIDs[hubID] = true
	}
	for _, f := range files {
		grp := categoryIndex[f.Category]
		hub := hubPos[f.Category]
		spread := math.Min(430, 60+30*math.Cbrt(float64(catCount[f.Category])))
		ox, oy, oz := hashOffset(f.Slug)
		ageDays := now.Sub(f.Date).Hours() / 24
		if ageDays < 0 {
			ageDays = 0
		}
		heat := math.Exp(-ageDays / 120)
		val := 1.2 + math.Min(3, float64(f.WordCount)/900) + math.Min(6, float64(degree[f.Slug]))
		nodes = append(nodes, GraphNode{
			ID: f.Slug, Name: f.Name, Val: val, Group: grp, Category: f.Category,
			FX: hub[0] + ox*spread, FY: hub[1] + oy*spread, FZ: hub[2] + oz*spread,
			Heat: heat, Deg: degree[f.Slug],
		})
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

// fibonacciSphere returns the i-th of n evenly distributed unit vectors.
func fibonacciSphere(i, n int) (x, y, z float64) {
	if n <= 1 {
		return 0, 1, 0
	}
	y = 1 - 2*(float64(i)+0.5)/float64(n)
	r := math.Sqrt(1 - y*y)
	phi := float64(i) * 2.399963229728653 // golden angle
	return math.Cos(phi) * r, y, math.Sin(phi) * r
}

// hashOffset derives a deterministic offset in [-1,1]^3 from a slug, biased
// toward the cluster core so clusters read as dense centers with sparse halos.
func hashOffset(s string) (x, y, z float64) {
	h := fnv.New64a()
	h.Write([]byte(s)) //nolint:errcheck
	v := h.Sum64()
	f := func(bits uint64) float64 {
		u := float64(bits&0xFFFFF) / float64(0xFFFFF) // [0,1]
		return (u*2 - 1)
	}
	x, y, z = f(v), f(v>>20), f(v>>40)
	// bias toward center
	n := math.Sqrt(x*x+y*y+z*z) + 1e-9
	scale := math.Pow(n/math.Sqrt(3), 0.5) / n * math.Sqrt(3)
	return x * scale, y * scale, z * scale
}

// GraphSubset keeps all hubs plus the top-N notes ranked by heat and
// connectivity, dropping links whose endpoints were pruned.
func (g GraphData) Subset(maxNotes int) GraphData {
	notes := 0
	for _, n := range g.Nodes {
		if !n.IsHub {
			notes++
		}
	}
	if maxNotes <= 0 || notes <= maxNotes {
		return g
	}
	var out GraphData
	byCat := map[string][]GraphNode{}
	for _, n := range g.Nodes {
		if n.IsHub {
			out.Nodes = append(out.Nodes, n)
		} else {
			byCat[n.Category] = append(byCat[n.Category], n)
		}
	}
	keep := map[string]bool{}
	for _, n := range out.Nodes {
		keep[n.ID] = true
	}
	// Proportional quota per category, filled by evenly-spaced sampling of
	// the category's recency-sorted notes — clusters keep their relative
	// size and the full hot→cold age spread stays visible.
	cats := make([]string, 0, len(byCat))
	for c := range byCat {
		cats = append(cats, c)
	}
	sort.Strings(cats)
	for _, c := range cats {
		grp := byCat[c]
		sort.Slice(grp, func(i, j int) bool {
			if grp[i].Heat != grp[j].Heat {
				return grp[i].Heat > grp[j].Heat
			}
			return grp[i].ID < grp[j].ID
		})
		quota := int(float64(maxNotes)*float64(len(grp))/float64(notes) + 0.5)
		if quota < 1 {
			quota = 1
		}
		if quota > len(grp) {
			quota = len(grp)
		}
		step := float64(len(grp)) / float64(quota)
		for k := 0; k < quota; k++ {
			n := grp[int(float64(k)*step)]
			out.Nodes = append(out.Nodes, n)
			keep[n.ID] = true
		}
	}
	for _, l := range g.Links {
		if keep[l.Source] && keep[l.Target] {
			out.Links = append(out.Links, l)
		}
	}
	return out
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

package service

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

type VaultFile struct {
	Slug     string
	Name     string
	Category string
	FullPath string
}

type SidebarItem struct {
	Slug string
	Name string
}

type SidebarCategory struct {
	Name  string
	Items []SidebarItem
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
	Files     []VaultFile
	Sidebar   []SidebarCategory
	GraphData GraphData
}

var wikilinkRe = regexp.MustCompile(`\[\[([^\]|]+)(?:\|[^\]]+)?\]\]`)

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	return s
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
		seen[slug] = true

		rel := strings.TrimPrefix(path, "raw/")
		parts := strings.SplitN(rel, "/", 2)
		category := "Uncategorized"
		if len(parts) > 1 {
			category = parts[0]
		}

		files = append(files, VaultFile{Slug: slug, Name: baseName, Category: category, FullPath: path})
		return nil
	})

	// Build sidebar
	grouped := map[string][]SidebarItem{}
	for _, f := range files {
		grouped[f.Category] = append(grouped[f.Category], SidebarItem{Slug: f.Slug, Name: f.Name})
	}
	var sidebar []SidebarCategory
	for cat, items := range grouped {
		sidebar = append(sidebar, SidebarCategory{Name: cat, Items: items})
	}
	// stable sort
	for i := 0; i < len(sidebar)-1; i++ {
		for j := i + 1; j < len(sidebar); j++ {
			if sidebar[i].Name > sidebar[j].Name {
				sidebar[i], sidebar[j] = sidebar[j], sidebar[i]
			}
		}
	}

	// Build graph
	categoryIndex := map[string]int{}
	catFiles := map[string][]VaultFile{}
	nextGroup := 0
	for _, f := range files {
		if _, ok := categoryIndex[f.Category]; !ok {
			categoryIndex[f.Category] = nextGroup
			nextGroup++
		}
		catFiles[f.Category] = append(catFiles[f.Category], f)
	}

	nodeIDs := map[string]bool{}
	var nodes []GraphNode
	var links []GraphLink

	for cat, grp := range categoryIndex {
		hubID := "hub-" + slugify(cat)
		nodes = append(nodes, GraphNode{ID: hubID, Name: cat, Val: max(3, len(catFiles[cat])/2), Group: grp, Category: cat, IsHub: true})
		nodeIDs[hubID] = true
	}
	for _, f := range files {
		grp := categoryIndex[f.Category]
		nodes = append(nodes, GraphNode{ID: f.Slug, Name: f.Name, Val: 1, Group: grp, Category: f.Category})
		nodeIDs[f.Slug] = true
		links = append(links, GraphLink{Source: f.Slug, Target: "hub-" + slugify(f.Category)})
	}

	// Wikilink edges
	for _, f := range files {
		data, err := fs.ReadFile(vaultFS, f.FullPath)
		if err != nil {
			continue
		}
		for _, m := range wikilinkRe.FindAllSubmatch(data, -1) {
			target := slugify(string(m[1]))
			if nodeIDs[target] && target != f.Slug {
				links = append(links, GraphLink{Source: f.Slug, Target: target})
			}
		}
	}

	return BookmarkIndex{Files: files, Sidebar: sidebar, GraphData: GraphData{Nodes: nodes, Links: links}}
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

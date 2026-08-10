package handler

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"myapp/internal/db"
	"myapp/internal/service"
	"myapp/ui/layouts"
)

const siteBase = "https://facorreia.com"

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	_, _ = w.Write([]byte(
		"User-agent: *\n" +
			"Allow: /\n" +
			"Disallow: /admin/\n" +
			"Disallow: /api/\n" +
			"Disallow: /book-call\n" +
			"\n" +
			"Sitemap: " + siteBase + "/sitemap.xml\n",
	))
}

type sitemapURL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty"`
}

type urlSet struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

type SitemapHandler struct {
	Q     *db.Queries
	Store *service.BookmarkStore
}

func (h *SitemapHandler) Serve(w http.ResponseWriter, r *http.Request) {
	set := urlSet{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	today := time.Now().UTC().Format("2006-01-02")

	type staticPage struct {
		path     string
		freq     string
		priority float32
	}
	statics := []staticPage{
		{"/", "weekly", 1.0},
		{"/projects", "weekly", 0.9},
		{"/about", "monthly", 0.8},
		{"/curriculum", "monthly", 0.8},
		{"/stack", "monthly", 0.6},
		{"/blog", "weekly", 0.8},
		{"/bookmarks", "weekly", 0.6},
		{"/proposal", "monthly", 0.7},
		{"/stats", "daily", 0.4},
	}
	for _, p := range statics {
		set.URLs = append(set.URLs, sitemapURL{
			Loc: siteBase + p.path, LastMod: today, ChangeFreq: p.freq, Priority: p.priority,
		})
	}
	for _, p := range allProjects {
		set.URLs = append(set.URLs, sitemapURL{
			Loc: siteBase + "/projects/" + p.Slug, LastMod: today, ChangeFreq: "monthly", Priority: 0.7,
		})
	}
	if h.Q != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		if posts, err := h.Q.ListBlogPosts(ctx); err == nil {
			for _, b := range posts {
				lastMod := today
				// Prefer SQLite updated_at, then the editorial date field.
				if lm := layouts.ParseBlogLastMod(b.UpdatedAt, b.Date); lm != "" {
					lastMod = lm
				}
				set.URLs = append(set.URLs, sitemapURL{
					Loc:        siteBase + "/blog/" + b.Slug,
					LastMod:    lastMod,
					ChangeFreq: "monthly",
					Priority:   0.6,
				})
			}
		}
	}
	if h.Store != nil {
		idx := h.Store.Index()
		for i := range idx.Files {
			f := &idx.Files[i]
			set.URLs = append(set.URLs, sitemapURL{
				Loc: siteBase + "/bookmarks/" + f.Slug, LastMod: f.Date.Format("2006-01-02"),
				ChangeFreq: "monthly", Priority: 0.4,
			})
		}
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_, _ = w.Write([]byte(xml.Header))
	_ = xml.NewEncoder(w).Encode(set)
}

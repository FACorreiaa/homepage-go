package handler

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"myapp/internal/db"
	"myapp/internal/service"
)

const siteBase = "https://facorreia.com"

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	_, _ = w.Write([]byte("User-agent: *\nAllow: /\nDisallow: /admin/\n\nSitemap: " + siteBase + "/sitemap.xml\n"))
}

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
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
	for _, p := range []string{"/", "/projects", "/about", "/curriculum", "/stack", "/blog", "/bookmarks", "/proposal", "/play"} {
		set.URLs = append(set.URLs, sitemapURL{Loc: siteBase + p})
	}
	for _, p := range allProjects {
		set.URLs = append(set.URLs, sitemapURL{Loc: siteBase + "/projects/" + p.Slug})
	}
	if h.Q != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		if posts, err := h.Q.ListBlogPosts(ctx); err == nil {
			for _, b := range posts {
				set.URLs = append(set.URLs, sitemapURL{Loc: siteBase + "/blog/" + b.Slug})
			}
		}
	}
	if h.Store != nil {
		idx := h.Store.Index()
		for i := range idx.Files {
			f := &idx.Files[i]
			set.URLs = append(set.URLs, sitemapURL{Loc: siteBase + "/bookmarks/" + f.Slug, LastMod: f.Date.Format("2006-01-02")})
		}
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_, _ = w.Write([]byte(xml.Header))
	_ = xml.NewEncoder(w).Encode(set)
}

package handler

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/yuin/goldmark"
	"myapp/internal/service"
	"myapp/ui/pages"
)

const bookmarksPerPage = 30

type BookmarksHandler struct {
	Store   *service.BookmarkStore
	VaultFS fs.FS
}

func (h *BookmarksHandler) List(w http.ResponseWriter, r *http.Request) {
	idx := h.Store.Index()

	q := r.URL.Query().Get("q")
	category := r.URL.Query().Get("category")
	sortKey := r.URL.Query().Get("sort")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	items, total := idx.Filter(q, category, sortKey, page, bookmarksPerPage)
	params := pages.BookmarkQuery{Q: q, Category: category, Sort: sortKey, Page: page, Total: total, PerPage: bookmarksPerPage}

	isHTMX := r.Header.Get("HX-Request") == "true"
	var err error
	switch {
	case isHTMX && page > 1:
		// "Load more": rows + replacement sentinel only.
		err = pages.BookmarkRows(items, params).Render(r.Context(), w)
	case isHTMX:
		// Search/filter/sort change: swap the whole results region.
		err = pages.BookmarkResults(items, params).Render(r.Context(), w)
	default:
		err = pages.Bookmarks(idx, items, params).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *BookmarksHandler) Show(w http.ResponseWriter, r *http.Request) {
	idx := h.Store.Index()
	slug := r.PathValue("slug")
	found := idx.BySlug[slug]
	if found == nil {
		http.NotFound(w, r)
		return
	}

	raw, err := fs.ReadFile(h.VaultFS, found.FullPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	cleaned := service.StripFrontmatter(string(raw))
	resolved := service.ResolveWikilinks(cleaned)

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(resolved), &buf); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
		return
	}

	related := idx.Related(found, 5)
	if err := pages.BookmarkPost(*found, buf.String(), related).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *BookmarksHandler) GraphData(w http.ResponseWriter, r *http.Request) {
	idx := h.Store.Index()
	limit := 2000
	if v, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil && v > 0 && v <= 10000 {
		limit = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=300")
	json.NewEncoder(w).Encode(idx.GraphData.Subset(limit)) //nolint:errcheck
}

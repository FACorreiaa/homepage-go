package handler

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/yuin/goldmark"
	"myapp/internal/service"
	"myapp/ui/pages"
)

type BookmarksHandler struct {
	Index   service.BookmarkIndex
	VaultFS fs.FS
}

func (h *BookmarksHandler) List(w http.ResponseWriter, r *http.Request) {
	if err := pages.Bookmarks(h.Index.Sidebar, len(h.Index.Files)).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *BookmarksHandler) Show(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	var found *service.VaultFile
	for i := range h.Index.Files {
		if h.Index.Files[i].Slug == slug {
			found = &h.Index.Files[i]
			break
		}
	}
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

	isHTMX := r.Header.Get("HX-Request") == "true"
	if err := pages.BookmarkPost(found.Name, found.Category, buf.String(), h.Index.Sidebar, slug, isHTMX).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *BookmarksHandler) GraphData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Index.GraphData) //nolint:errcheck
}

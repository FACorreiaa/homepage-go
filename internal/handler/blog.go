package handler

import (
	"bytes"
	"net/http"

	"github.com/yuin/goldmark"
	"myapp/internal/db"
	"myapp/internal/service"
	"myapp/ui/pages"
)

type BlogHandler struct {
	Q       *db.Queries
	Tracker *service.BlogTracker
}

func (h *BlogHandler) List(w http.ResponseWriter, r *http.Request) {
	h.Tracker.RecordVisit(r)

	posts, err := h.Q.ListBlogPosts(r.Context())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	pages.BlogList(posts, h.Tracker.Snapshot()).Render(r.Context(), w)
}

func (h *BlogHandler) Post(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := h.Q.GetBlogPostBySlug(r.Context(), slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(post.Body), &buf); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
		return
	}
	pages.BlogPost(post, buf.String()).Render(r.Context(), w)
}

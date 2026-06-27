package handler

import (
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"myapp/internal/db"
	sess "myapp/internal/session"
	"myapp/ui/pages"
)

type AdminHandler struct {
	Q *db.Queries
}

func (h *AdminHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.Login("").Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) LoginSubmit(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.FormValue("email"))
	password := r.FormValue("password")

	user, err := h.Q.GetUserByEmail(r.Context(), email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		if err := pages.Login("Invalid credentials.").Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
		return
	}
	sess.SetUserID(w, r, user.ID) //nolint:errcheck
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	leads, err := h.Q.ListProposalLeads(r.Context())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if err := pages.Dashboard(leads).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) Logout(w http.ResponseWriter, r *http.Request) {
	sess.Clear(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AdminHandler) BlogNewPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.BlogEditor(db.BlogPost{}, false).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) BlogCreate(w http.ResponseWriter, r *http.Request) {
	params := db.CreateBlogPostParams{
		Slug:     strings.TrimSpace(r.FormValue("slug")),
		Title:    strings.TrimSpace(r.FormValue("title")),
		Summary:  strings.TrimSpace(r.FormValue("summary")),
		Category: strings.TrimSpace(r.FormValue("category")),
		Date:     strings.TrimSpace(r.FormValue("date")),
		Body:     r.FormValue("body"),
	}
	if err := h.Q.CreateBlogPost(r.Context(), params); err != nil {
		if err := pages.BlogEditor(db.BlogPost{Slug: params.Slug, Title: params.Title, Summary: params.Summary, Category: params.Category, Date: params.Date, Body: params.Body}, false).Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, r, "/blog", http.StatusSeeOther)
}

func (h *AdminHandler) BlogEditPage(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := h.Q.GetBlogPostBySlug(r.Context(), slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if err := pages.BlogEditor(post, true).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) BlogUpdate(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	params := db.UpdateBlogPostParams{
		Title:    strings.TrimSpace(r.FormValue("title")),
		Summary:  strings.TrimSpace(r.FormValue("summary")),
		Category: strings.TrimSpace(r.FormValue("category")),
		Date:     strings.TrimSpace(r.FormValue("date")),
		Body:     r.FormValue("body"),
		Slug:     slug,
	}
	if err := h.Q.UpdateBlogPost(r.Context(), params); err != nil {
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/blog", http.StatusSeeOther)
}

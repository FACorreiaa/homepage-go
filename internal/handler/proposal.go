package handler

import (
	"database/sql"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
	htmx "github.com/angelofallars/htmx-go"
	"myapp/internal/db"
	"myapp/internal/service"
	"myapp/ui/pages"
)

var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type ProposalHandler struct {
	Q *db.Queries
}

func (h *ProposalHandler) Form(w http.ResponseWriter, r *http.Request) {
	pages.Proposal("", "").Render(r.Context(), w)
}

func (h *ProposalHandler) Submit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Honeypot
	if r.FormValue("website") != "" {
		pages.ProposalSuccess().Render(r.Context(), w)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))
	projectType := r.FormValue("project_type")
	details := strings.TrimSpace(r.FormValue("details"))

	if name == "" || email == "" || projectType == "" || details == "" {
		if htmx.IsHTMX(r) {
			pages.ProposalResult("", "Please fill out all required fields.").Render(r.Context(), w)
		} else {
			pages.Proposal("", "Please fill out all required fields.").Render(r.Context(), w)
		}
		return
	}
	if !emailRe.MatchString(email) {
		if htmx.IsHTMX(r) {
			pages.ProposalResult("", "Please enter a valid email address.").Render(r.Context(), w)
		} else {
			pages.Proposal("", "Please enter a valid email address.").Render(r.Context(), w)
		}
		return
	}

	lead := db.CreateProposalLeadParams{
		ID:          uuid.NewString(),
		Name:        name,
		Email:       email,
		Company:     sql.NullString{String: r.FormValue("company"), Valid: r.FormValue("company") != ""},
		ProjectType: projectType,
		Budget:      sql.NullString{String: r.FormValue("budget"), Valid: r.FormValue("budget") != ""},
		Timeline:    sql.NullString{String: r.FormValue("timeline"), Valid: r.FormValue("timeline") != ""},
		Details:     details,
	}
	if err := h.Q.CreateProposalLead(r.Context(), lead); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	service.SendLeadNotification(os.Getenv("DISCORD_WEBHOOK_URL"), db.ProposalLead{
		Name: name, Email: email, ProjectType: projectType, Details: details,
	})

	if htmx.IsHTMX(r) {
		pages.ProposalSuccess().Render(r.Context(), w)
	} else {
		pages.Proposal("success", "").Render(r.Context(), w)
	}
}

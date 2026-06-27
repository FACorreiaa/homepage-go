package handler

import (
	"net/http"
	"os"
	"strings"

	"myapp/ui/pages"
)

func Play(w http.ResponseWriter, r *http.Request) {
	if err := pages.Play().Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func BookCall(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, normalizeCalendly(os.Getenv("CALENDLY_URL")), http.StatusTemporaryRedirect)
}

func normalizeCalendly(raw string) string {
	const defaultURL = "https://calendly.com/fernandocorreia316"
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return defaultURL
	}
	if strings.HasPrefix(raw, "https://") || strings.HasPrefix(raw, "http://") {
		return raw
	}
	if strings.HasPrefix(raw, "calendly.com/") {
		return "https://" + raw
	}
	return "https://calendly.com/" + strings.Trim(raw, "/")
}

package handler

import (
	"net/http"

	"myapp/ui/pages"
)

func Stack(w http.ResponseWriter, r *http.Request) {
	if err := pages.Stack().Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

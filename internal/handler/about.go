package handler

import (
	"net/http"

	"myapp/ui/pages"
)

func About(w http.ResponseWriter, r *http.Request) {
	if err := pages.About().Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

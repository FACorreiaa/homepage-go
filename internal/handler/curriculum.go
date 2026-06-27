package handler

import (
	"net/http"

	"myapp/ui/pages"
)

func Curriculum(w http.ResponseWriter, r *http.Request) {
	if err := pages.Curriculum().Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

package handler

import (
	"net/http"

	"myapp/ui/pages"
)

func About(w http.ResponseWriter, r *http.Request) {
	pages.About().Render(r.Context(), w)
}

package handler

import (
	"net/http"

	"myapp/ui/pages"
)

func Stack(w http.ResponseWriter, r *http.Request) {
	pages.Stack().Render(r.Context(), w)
}

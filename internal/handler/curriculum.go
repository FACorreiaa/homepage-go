package handler

import (
	"net/http"

	"myapp/ui/pages"
)

func Curriculum(w http.ResponseWriter, r *http.Request) {
	pages.Curriculum().Render(r.Context(), w)
}

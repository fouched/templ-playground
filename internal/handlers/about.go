package handlers

import (
	"net/http"

	"github.com/fouched/templ-playground/internal/render"
	"github.com/fouched/templ-playground/templates"
)

func About(w http.ResponseWriter, r *http.Request) {

	component := templates.About()
	_ = render.Template(w, r, component)
}

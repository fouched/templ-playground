package handlers

import (
	"net/http"

	"github.com/fouched/templ-playground/internal/render"
	"github.com/fouched/templ-playground/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {

	component := templates.Home("Fouche du Preez")
	_ = render.Template(w, r, component)
}

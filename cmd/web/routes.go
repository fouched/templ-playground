package main

import (
	"net/http"

	"github.com/fouched/templ-playground/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}
package http

import (
	"net/http"

	"github.com/KhachikAstoyan/go-api-template/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitHandlers(app *core.App) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	initUserHandlers(app, r)
	http.ListenAndServe(":8080", r)
}

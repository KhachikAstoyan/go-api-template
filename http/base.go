package http

import (
	"net/http"

	"github.com/KhachikAstoyan/go-api-template/core"
	"github.com/KhachikAstoyan/go-api-template/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitHandlers(app *core.App) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userService := services.GetUserService(app.DB)
	initUserHandlers(app, r, userService)

	http.ListenAndServe(":8080", r)
}

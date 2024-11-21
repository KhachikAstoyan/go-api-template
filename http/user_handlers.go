package http

import (
	"encoding/json"
	"net/http"

	"github.com/KhachikAstoyan/go-api-template/core"
	"github.com/KhachikAstoyan/go-api-template/services"
	"github.com/go-chi/chi/v5"
)

func initUserHandlers(app *core.App, mux *chi.Mux) {
	r := chi.NewRouter()
	mux.Mount("/user", r)

	c := userController{
		app:     app,
		service: services.GetUserService(app.DB),
	}

	r.Get("/", c.count)

}

type userController struct {
	app     *core.App
	service services.UserService
}

func (c *userController) count(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	count, err := c.service.CountUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(count)
}

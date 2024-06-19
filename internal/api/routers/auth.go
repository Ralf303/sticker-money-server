package routers

import (
	"example.com/stickerMoneyAdmin/internal/api/services"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/reg", services.RegisterUser)
	router.Get("/login", services.LoginUser)

	return router
}

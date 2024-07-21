package routers

import (
	"example.com/stickerMoneyAdmin/internal/api/services"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middleware.AuthMiddleware)
	router.Get("/getAll", services.GetUsers)
	router.Get("/getOne", services.GetUser)
	router.Put("/updateBalance", services.UpdateBalance)
	router.Put("/updateBan", services.UpdateBan)
	router.Get("/count", services.CountUsers)

	return router
}

package routers

import (
	"example.com/stickerMoneyAdmin/internal/api/services"
	"github.com/go-chi/chi/v5"
)

func StakeRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/getStakes", services.GetStakes)
	router.Put("/updateStake", services.UpdateStake)

	return router
}

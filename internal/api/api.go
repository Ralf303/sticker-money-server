package api

import (
	"example.com/stickerMoneyAdmin/internal/api/routers"
	"example.com/stickerMoneyAdmin/internal/api/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Mount("/users", routers.UserRoutes())
	router.Mount("/stakes", routers.StakeRoutes())
	router.Mount("/auth", routers.AuthRoutes())
	router.Get("/ping", services.Ping)
	return router
}

package router

import (
	"event-planner/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New(h handlers.Handlers) http.Handler {

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		}),
		middleware.Recoverer,
	)

	return router
}

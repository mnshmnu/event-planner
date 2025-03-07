package router

import (
	"event-planner/internal/handlers"
	"event-planner/pkg/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New(h handlers.Handlers) http.Handler {

	r := chi.NewRouter()

	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		}),
		middleware.Recoverer,
	)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
	})

	r.With(middlewares.JWTAuthenticate()).Route("/", func(r chi.Router) {
		r.Route("/user/availability", func(r chi.Router) {
			r.Post("/add", h.AddAvailability)
			r.Post("/get", h.GetAvailableSlotsHandler)
		})

		r.Route("/meeting", func(r chi.Router) {
			r.Post("/add", h.ScheduleMeeting)
			r.Post("/get", h.GetAvailableSlotsHandler)
		})
	})

	return r
}

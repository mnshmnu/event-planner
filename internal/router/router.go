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
			r.Get("/", h.GetAvailability) // id
			r.Post("/add", h.CreateAvailability)
			r.Patch("/", h.UpdateAvailability)  // id
			r.Delete("/", h.DeleteAvailability) // id
		})

		r.Route("/event", func(r chi.Router) {
			r.Get("/", h.GetEventByID) // id
			r.Post("/add", h.CreateEvent)
			r.Delete("/add", h.DeleteEvent) //id
			r.Patch("/add", h.UpdateEvent)
			r.Patch("/confirm", h.ConfirmFinalSlot)
		})
	})

	return r
}

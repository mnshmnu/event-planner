package handlers

import (
	"event-planner/internal/services"
	"net/http"
)

type Handlers interface {

	// user handlers

	// auth
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	// availability
	CreateAvailability(w http.ResponseWriter, r *http.Request)
	GetAvailability(w http.ResponseWriter, r *http.Request)
	UpdateAvailability(w http.ResponseWriter, r *http.Request)
	DeleteAvailability(w http.ResponseWriter, r *http.Request)

	// meeting
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	GetAllEventsByUser(w http.ResponseWriter, r *http.Request)
	ConfirmFinalSlot(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service services.Service
}

func New(s services.Service) Handlers {
	return &handler{
		service: s,
	}
}

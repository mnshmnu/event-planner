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
	AddAvailability(w http.ResponseWriter, r *http.Request)
	GetAvailableSlots(w http.ResponseWriter, r *http.Request)

	// meeting
	GetMeetings(w http.ResponseWriter, r *http.Request)
	ScheduleMeeting(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service services.Service
}

func New(s services.Service) Handlers {
	return &handler{
		service: s,
	}
}

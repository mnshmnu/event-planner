package handlers

import (
	"event-planner/internal/services"
	"net/http"
)

type Handlers interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service services.Service
}

func New(s services.Service) Handlers {
	return &handler{
		service: s,
	}
}

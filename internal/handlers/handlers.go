package handlers

import (
	"event-planner/internal/services"
)

type Handlers interface{}

type handler struct {
	service services.Service
}

func New(s services.Service) Handlers {
	return &handler{
		service: s,
	}
}

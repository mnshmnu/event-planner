package services

import "event-planner/internal/models"

type Service interface{}

type service struct {
	models models.Model
}

func New(m models.Model) Service {
	return &service{
		models: m,
	}
}

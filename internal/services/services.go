package services

import (
	"context"
	"event-planner/internal/entities"
	"event-planner/internal/models"
)

type Service interface {
	RegisterUser(ctx context.Context, user *entities.User) error
	AuthenticateUser(ctx context.Context, email, password string) (string, error)
}

type service struct {
	models models.Model
}

func New(m models.Model) Service {
	return &service{
		models: m,
	}
}

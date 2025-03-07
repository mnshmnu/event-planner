package services

import (
	"context"
	"event-planner/internal/entities"
	"event-planner/internal/models"
	"time"
)

type Service interface {
	RegisterUser(ctx context.Context, user *entities.User) error
	AuthenticateUser(ctx context.Context, email, password string) (string, error)

	AddAvailability(ctx context.Context, availability *entities.Availability) error
	GetAvailableSlots(ctx context.Context, duration int) (map[time.Time][]int, error)

	ScheduleMeeting(ctx context.Context, organizerID int, duration int) (*entities.Meeting, error)
}

type service struct {
	models models.Model
}

func New(m models.Model) Service {
	return &service{
		models: m,
	}
}

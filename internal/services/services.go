package services

import (
	"context"
	"event-planner/internal/entities"
	"event-planner/internal/entities/packet"
	"event-planner/internal/models"
	"event-planner/pkg/auth"
)

type Service interface {
	RegisterUser(ctx context.Context, user *entities.User) error
	AuthenticateUser(ctx context.Context, email, password string) (string, error)

	CreateAvailability(ctx context.Context, a *entities.ParticipantAvailability) (int64, error)
	GetAvailabilityByID(ctx context.Context, id int64) (*entities.ParticipantAvailability, error)
	UpdateAvailability(ctx context.Context, a *entities.ParticipantAvailability) error
	DeleteAvailability(ctx context.Context, id int64) error
	GetAvailabilitiesByEvent(ctx context.Context, eventID int64) ([]entities.ParticipantAvailability, error)

	CreateEventWithSlots(ctx context.Context, req packet.CreateEventReq) (*packet.CreateEventResp, error)
	GetEventByID(ctx context.Context, eventID int64) (*entities.Event, error)
	UpdateEvent(ctx context.Context, event *entities.Event) error
	DeleteEvent(ctx context.Context, eventID int64) error
	GetEventsByUser(ctx context.Context, userID int64) ([]*entities.Event, error)
	ConfirmFinalSlot(ctx context.Context, req packet.ConfirmSlotReq) error
}

type service struct {
	models models.Model
	auth   auth.Auth
}

func New(m models.Model, auth auth.Auth) Service {
	return &service{
		models: m,
		auth:   auth,
	}
}

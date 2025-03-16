package models

import (
	"context"
	"event-planner/internal/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type model struct {
	db *pgxpool.Pool
}

type Model interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, string, error)
	CreateUser(ctx context.Context, user *entities.User, hPass string) error

	CreateEvent(ctx context.Context, event *entities.Event) (int64, error)
	GetEventByID(ctx context.Context, id int64) (*entities.Event, error)
	UpdateEvent(ctx context.Context, event *entities.Event) error
	DeleteEvent(ctx context.Context, id int64) error
	GetAllEventsByUser(ctx context.Context, userID int64) ([]*entities.Event, error)

	CreateEventSlot(ctx context.Context, slot *entities.EventSlot) (int64, error)
	CreateEventSlots(ctx context.Context, slots []entities.EventSlot) error
	IsSlotPartOfEvent(ctx context.Context, slotID, eventID int64) (bool, error)
	ConfirmFinalSlot(ctx context.Context, eventID, slotID int64) error

	CreateAvailability(ctx context.Context, availability *entities.ParticipantAvailability) (int64, error)
	GetAvailabilityByID(ctx context.Context, id int64) (*entities.ParticipantAvailability, error)
	UpdateAvailability(ctx context.Context, availability *entities.ParticipantAvailability) error
	DeleteAvailability(ctx context.Context, id int64) error
	GetAvailabilitiesByEvent(ctx context.Context, eventID int64) ([]entities.ParticipantAvailability, error)
}

func New(db *pgxpool.Pool) Model {
	return &model{
		db: db,
	}
}

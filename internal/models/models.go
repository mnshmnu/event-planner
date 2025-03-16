package models

import (
	"context"
	"event-planner/internal/entities"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type model struct {
	db *pgxpool.Pool
}

type Model interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, string, error)
	CreateUser(ctx context.Context, user *entities.User, hPass string) error

	GetMeetings(ctx context.Context) ([]*entities.Meeting, error)
	CreateMeeting(ctx context.Context, meeting *entities.Meeting) error

	AddAvailability(ctx context.Context, availability *entities.Availability) error
	GetAvailabilities(ctx context.Context, duration int) (map[time.Time][]int, error)
}

func New(db *pgxpool.Pool) Model {
	return &model{
		db: db,
	}
}

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
}

func New(db *pgxpool.Pool) Model {
	return &model{
		db: db,
	}
}

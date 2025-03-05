package models

import "github.com/jackc/pgx/v5/pgxpool"

type model struct {
	db *pgxpool.Pool
}

type Model interface{}

func New(db *pgxpool.Pool) Model {
	return &model{
		db: db,
	}
}

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitTables(db *pgxpool.Pool) error {
	const initSQL = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		duration INT NOT NULL
	);
	`

	ctx := context.Background()
	_, err := db.Exec(ctx, initSQL)
	return err
}

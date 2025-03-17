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
		title TEXT NOT NULL,
		created_by INTEGER REFERENCES users(id),
		duration INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		confirmed_slot_id BIGINT REFERENCES event_slots(id)
	);

	CREATE TABLE IF NOT EXISTS event_slots (
		id SERIAL PRIMARY KEY,
		event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL
	);

	CREATE TABLE IF NOT EXISTS event_participants (
		id SERIAL PRIMARY KEY,
		event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS participant_availabilities (
		id SERIAL PRIMARY KEY,
		event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL
	);
	`

	ctx := context.Background()
	_, err := db.Exec(ctx, initSQL)
	return err
}

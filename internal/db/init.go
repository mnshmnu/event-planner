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

	CREATE TABLE IF NOT EXISTS availabilities (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS meetings (
		id SERIAL PRIMARY KEY,
		organizer_id INT NOT NULL,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL,
		FOREIGN KEY (organizer_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS meeting_participants (
		meeting_id INT NOT NULL,
		user_id INT NOT NULL,
		PRIMARY KEY (meeting_id, user_id),
		FOREIGN KEY (meeting_id) REFERENCES meetings(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`

	ctx := context.Background()
	_, err := db.Exec(ctx, initSQL)
	return err
}

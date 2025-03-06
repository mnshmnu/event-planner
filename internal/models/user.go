package models

import (
	"context"
	"database/sql"
	"event-planner/internal/entities"
	"fmt"
)

func (m *model) GetUserByEmail(ctx context.Context, email string) (*entities.User, string, error) {
	var user entities.User
	var hashedPassword string
	err := m.db.QueryRow(ctx, "SELECT id, name, email, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", fmt.Errorf("user not found")
		}
		return nil, "", err
	}

	return &user, hashedPassword, nil
}

func (m *model) CreateUser(ctx context.Context, user *entities.User, hPass string) error {

	_, err := m.db.
		Exec(ctx,
			"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
			user.Name, user.Email, string(hPass),
		)
	return err
}

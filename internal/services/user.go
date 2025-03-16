package services

import (
	"context"
	"database/sql"
	"errors"
	"event-planner/internal/entities"
	"fmt"
)

func (s *service) RegisterUser(ctx context.Context, user *entities.User) error {
	hPass, err := s.auth.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	err = s.models.CreateUser(ctx, user, string(hPass))
	if err != nil {
		return err
	}

	return err
}

func (s *service) AuthenticateUser(ctx context.Context, email, password string) (string, error) {
	var user *entities.User
	var hashedPassword string
	user, hashedPassword, err := s.models.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	isValid, _ := s.auth.CompareHash([]byte(hashedPassword), []byte(password))
	if !isValid {
		return "", errors.New("invalid credentials")
	}

	token, err := s.auth.GenerateJWTToken(map[string]interface{}{
		"userID": user.ID,
		"email":  user.Email,
		"name":   user.Name,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

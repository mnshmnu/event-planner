package services

import (
	"context"
	"database/sql"
	"errors"
	"event-planner/internal/entities"
	"event-planner/pkg/auth"
	"event-planner/pkg/middlewares"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUser(ctx context.Context, user *entities.User) error {
	hPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
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

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWTToken(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	})

	return token, nil
}

func (h *service) AddAvailability(ctx context.Context, availability *entities.Availability) error {

	curUser := middlewares.GetCurrentUser(ctx)

	zap.S().Infow("Current User", "user", curUser.Email, "ctx", ctx)
	return nil
}

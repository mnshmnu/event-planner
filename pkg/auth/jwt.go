package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(payload map[string]any) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}
	for k, v := range payload {
		claims[k] = v
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}

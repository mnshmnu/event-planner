package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"event-planner/internal/entities"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func returnResp(w http.ResponseWriter, message string, statusCode int, err error) {
	zap.S().Debugw(message, "err", err)

	res, _ := json.Marshal(map[string]string{
		"message": message,
	})
	w.WriteHeader(statusCode)
	w.Write(res)
}

type contextKey string

var JWTContextKey = contextKey(os.Getenv("JWT_CONTEXT_KEY"))

func JWTAuthenticate() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtHeader := r.Header["Authorization"]
			if len(jwtHeader) == 0 {
				returnResp(w, "No JWT token provided", http.StatusUnauthorized, errors.New("no JWT token"))
				return
			}

			tokenStr := strings.Split(jwtHeader[0], " ")
			jwtSecret := os.Getenv("JWT_SECRET")
			if len(tokenStr) != 2 {
				returnResp(w, "Invalid JWT token provided", http.StatusUnauthorized, errors.New("invalid JWT token"))
				return
			}

			token, err := jwt.ParseWithClaims(tokenStr[1], &entities.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				returnResp(w, "Failed to parse JWT", http.StatusUnauthorized, err)
				return
			}

			claims, ok := token.Claims.(*entities.UserClaims)
			if !ok {
				returnResp(w, "Failed to parse JWT claims", http.StatusUnauthorized, errors.New("failed to parse JWT claims"))
				return
			}

			ctx := context.WithValue(r.Context(), JWTContextKey, claims)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// GetCurrentUser takes out user details from context and returns it
func GetCurrentUser(ctx context.Context) *entities.UserClaims {
	user := ctx.Value(JWTContextKey).(*entities.UserClaims)
	if user == nil {
		return nil
	}
	return user
}

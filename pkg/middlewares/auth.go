package middlewares

import (
	"context"
	"encoding/json"
	"errors"
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

func JWTAuthenticate() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtHeader := r.Header["Authorization"]
			if len(jwtHeader) == 0 {
				returnResp(w, "No JWT token provided", http.StatusUnauthorized, errors.New("no JWT token"))
				return
			}

			tokenStr := strings.Split(jwtHeader[0], " ")
			if len(tokenStr) != 2 {
				returnResp(w, "Invalid JWT token provided", http.StatusUnauthorized, errors.New("invalid JWT token"))
				return
			}
			zap.S().Infow("JWT token", "token", tokenStr[1])

			token, err := jwt.Parse(tokenStr[1], func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				returnResp(w, "Failed to parse JWT", http.StatusUnauthorized, err)
				return
			}

			jwtKey := contextKey(os.Getenv("JWT_CONTEXT_KEY"))

			ctx := context.WithValue(r.Context(), jwtKey, token)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

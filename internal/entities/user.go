package entities

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	UserID int    `json:"userID"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

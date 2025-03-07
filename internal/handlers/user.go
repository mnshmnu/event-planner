package handlers

import (
	"encoding/json"
	"event-planner/internal/entities"
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		zap.S().Debugw("Failed to decode request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.service.RegisterUser(r.Context(), &user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		zap.S().Debugw("Failed to decode request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	jwt, err := h.service.AuthenticateUser(r.Context(), credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"token": jwt})
}

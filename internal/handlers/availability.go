package handlers

import (
	"encoding/json"
	"event-planner/internal/entities"
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) GetAvailableSlotsHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Duration int `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	slots, err := h.service.GetAvailableSlots(r.Context(), request.Duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(slots)
}

func (h *handler) AddAvailability(w http.ResponseWriter, r *http.Request) {
	var availability entities.Availability

	if err := json.NewDecoder(r.Body).Decode(&availability); err != nil {
		zap.S().Debugw("Failed to decode request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.service.AddAvailability(r.Context(), &availability)
	if err != nil {
		http.Error(w, "Failed to add availability", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Availability added successfully",
	})
}

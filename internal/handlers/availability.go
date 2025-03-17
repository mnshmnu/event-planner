package handlers

import (
	"encoding/json"
	"event-planner/internal/entities"
	"event-planner/internal/entities/packet"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

// CreateAvailabilityHandler handles POST /availability
func (h *handler) CreateAvailability(w http.ResponseWriter, r *http.Request) {

	var input packet.CreateAvailabilityReq

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		zap.S().Debugw("Failed to decode request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	a := entities.ParticipantAvailability{
		UserID:    input.UserID,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
	}

	id, err := h.service.CreateAvailability(r.Context(), &a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"availability_id": id})
}

// GetAvailabilityHandler handles GET /availability/{id}
func (h *handler) GetAvailability(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	availability, err := h.service.GetAvailabilityByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Availability not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(availability)
}

// UpdateAvailabilityHandler handles PUT /availability/{id}
func (h *handler) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	var input packet.UpdateAvailabilityReq

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	availability, err := h.service.GetAvailabilityByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Availability not found", http.StatusNotFound)
		return
	}

	availability.StartTime = input.StartTime
	availability.EndTime = input.EndTime

	err = h.service.UpdateAvailability(r.Context(), availability)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

// DeleteAvailabilityHandler handles DELETE /availability/{id}
func (h *handler) DeleteAvailability(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteAvailability(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete availability", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

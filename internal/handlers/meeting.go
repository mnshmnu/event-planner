package handlers

import (
	"encoding/json"
	"event-planner/pkg/middlewares"
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) ScheduleMeeting(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Duration int `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.S().Debugw("Failed to decode request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userID := middlewares.GetCurrentUser(r.Context()).UserID

	meeting, err := h.service.ScheduleMeeting(r.Context(), userID, req.Duration)
	if err != nil {
		http.Error(w, "Failed to schedule meeting", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Meeting added successfully",
		"res":     meeting,
	})
}

func (h *handler) GetMeetings(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Duration int `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.S().Debugw("Failed to decode request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	meetList, err := h.service.GetMeetings(r.Context())
	if err != nil {
		http.Error(w, "Failed to schedule meeting", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Meeting added successfully",
		"res":     meetList,
	})
}

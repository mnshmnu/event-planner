package handlers

import (
	"encoding/json"
	"event-planner/internal/entities"
	"event-planner/internal/entities/packet"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input packet.CreateEventReq
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateEventWithSlots(r.Context(), input)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event id", http.StatusBadRequest)
		return
	}

	event, err := h.service.GetEventByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(event)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var input packet.UpdateEventReq

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		zap.S().Debugw("Failed to decode update event request body", "err", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	event := entities.Event{
		ID:              input.ID,
		Title:           input.Title,
		DurationMinutes: input.DurationMinutes,
	}

	err := h.service.UpdateEvent(r.Context(), &event)
	if err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteEvent(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

func (h *handler) GetAllEventsByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

func (h *handler) ConfirmFinalSlot(w http.ResponseWriter, r *http.Request) {
	var req packet.ConfirmSlotReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.ConfirmFinalSlot(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "Event slot is confirmed"})
}

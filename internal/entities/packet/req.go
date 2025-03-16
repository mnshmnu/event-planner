package packet

import (
	"event-planner/internal/entities"
	"time"
)

// availability
type UpdateAvailabilityReq struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type CreateAvailabilityReq struct {
	EventID   int64     `json:"eventID"`
	UserID    int64     `json:"userID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// event

type CreateEventReq struct {
	Title           string               `json:"title"`
	CreatedBy       int64                `json:"created_by"`
	DurationMinutes int                  `json:"duration_minutes"`
	Slots           []entities.EventSlot `json:"slots"`
}

type UpdateEventReq struct {
	ID              int64  `json:"id"`
	Title           string `json:"title"`
	DurationMinutes int    `json:"duration_minutes"`
}

type ConfirmSlotReq struct {
	EventID int64 `json:"event_id"`
	SlotID  int64 `json:"slot_id"`
}

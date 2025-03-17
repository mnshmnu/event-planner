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
	UserID    int64     `json:"userID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// event

type CreateEventReq struct {
	Title           string               `json:"title"`
	CreatedBy       int64                `json:"createdBy"`
	DurationMinutes int                  `json:"durationMinutes"`
	Slots           []entities.EventSlot `json:"slots"`
}

type UpdateEventReq struct {
	ID              int64  `json:"id"`
	Title           string `json:"title"`
	DurationMinutes int    `json:"durationMinutes"`
}

type ConfirmSlotReq struct {
	EventID int64 `json:"eventID"`
	SlotID  int64 `json:"slotID"`
}

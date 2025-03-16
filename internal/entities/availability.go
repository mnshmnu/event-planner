package entities

import "time"

type ParticipantAvailability struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"event_id"`
	UserID    int64     `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

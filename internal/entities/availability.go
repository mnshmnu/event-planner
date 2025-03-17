package entities

import "time"

type ParticipantAvailability struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

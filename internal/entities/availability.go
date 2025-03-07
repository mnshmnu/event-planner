package entities

import "time"

type Availability struct {
	UserID    int       `json:"userID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

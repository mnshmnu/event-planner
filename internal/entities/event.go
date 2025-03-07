package entities

import "time"

type Event struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Duration    int         `json:"duration"`
	Slots       []time.Time `json:"slots"`
}

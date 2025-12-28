package scheduler

import "time"

// --- Structures ---

type Agenda struct {
	ID      int    `json:"id"`
	GroupID string `json:"group_id"`
	ICalURL string `json:"ical_url"`
}

type Event struct {
	Summary string
	Start   time.Time
	End     time.Time
}

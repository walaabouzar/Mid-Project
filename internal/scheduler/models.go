package scheduler

import "time"

// --- Structures ---

type Agenda struct {
	ID      int    `json:"id"`
	GroupID string `json:"group_id"`
	ICalURL string `json:"ical_url"`
}

type Event struct {
    AgendaID int       `json:"agenda_id"`  
    GroupID  string    `json:"group_id"`   
    Summary  string    `json:"summary"`
    Start    time.Time `json:"start"`
    End      time.Time `json:"end"`
}


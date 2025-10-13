package models


import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

type Event struct {
	// Id est un UUID (pointeur) — peut être nil avant insertion
	ID *uuid.UUID `json:"id,omitempty" db:"id"`
	// Uid est l'identifiant fourni par l'API .ics (ex: "ADE60323...")
	UID string `json:"uid" db:"u_uid"`
	Title       string    `json:"title" db:"title"`
	Location    string    `json:"location" db:"location"`
	Description string    `json:"description" db:"description"`
	Start       time.Time `json:"start" db:"start"`
	End         time.Time `json:"end" db:"end"`
	AgendaID    int       `json:"agenda_id" db:"agenda_id"`
}
// Agenda représente un agenda spécifique
type Agenda struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

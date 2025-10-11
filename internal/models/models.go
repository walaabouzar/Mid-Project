package models
import "time"
import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}
// Agenda représente un emploi du temps (ex: M1 Groupe 2 langue)
type Agenda struct {
	ID          int       `json:"id" db:"id"`
	AgendaID    string    `json:"agenda_id" db:"agenda_id"`     // Ex: "13345"
	Name        string    `json:"name" db:"name"`               // Ex: "M1 Groupe 2 langue"
	Description string    `json:"description" db:"description"` // Description optionnelle
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Event représente un cours dans l'emploi du temps
type Event struct {
	ID           int       `json:"id" db:"id"`
	AgendaID     string    `json:"agenda_id" db:"agenda_id"`           // Lien vers l'agenda
	UID          string    `json:"uid" db:"uid"`                       // Identifiant unique du iCal
	Summary      string    `json:"summary" db:"summary"`               // Titre du cours
	Location     string    `json:"location" db:"location"`             // Salle
	Description  string    `json:"description" db:"description"`       // Description complète
	StartTime    time.Time `json:"start_time" db:"start_time"`         // Date/heure début
	EndTime      time.Time `json:"end_time" db:"end_time"`             // Date/heure fin
	LastModified time.Time `json:"last_modified" db:"last_modified"`   // Dernière modification (iCal)
	Sequence     int       `json:"sequence" db:"sequence"`             // Numéro de version (iCal)
	CreatedAt    time.Time `json:"created_at" db:"created_at"`         // Date création dans notre DB
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`         // Date mise à jour dans notre DB
}
// Liste des agendas disponibles
var Agendas = []Agenda{
	{AgendaID: "13295", Name: "M1 Groupe 1 langue"},
	{AgendaID: "13345", Name: "M1 Groupe 2 langue"},
	{AgendaID: "13397", Name: "M1 Groupe 3 langue"},
	{AgendaID: "7224", Name: "M1 Groupe 1 option"},
	{AgendaID: "7225", Name: "M1 Groupe 2 option"},
	{AgendaID: "62962", Name: "M1 Groupe 3 option"},
	{AgendaID: "62090", Name: "M1 Groupe option"},
	{AgendaID: "56529", Name: "M1 - Tutorat L2"},
}
package models

import "time"
import "github.com/gofrs/uuid"

type Event struct {
	ID          uuid.UUID `json:"id"`          // UUID
	AgendaID    int       `json:"agenda_id"`   // Agenda ID
	UID         string    `json:"uid"`         // Identifiant unique ADE
	Description string    `json:"description"` // Détails du cours
	Name        string    `json:"name"`        // Nom du cours
	Start       time.Time `json:"start"`       // Date/heure de début
	End         time.Time `json:"end"`         // Date/heure de fin
	Location    string    `json:"location"`    // Salle / lieu
	LastUpdate  time.Time `json:"lastUpdate"`  // Dernière mise à jour
}

package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

type Agenda struct {
	ID      int    `json:"id"`
	GroupID string `json:"group_id"`
	ICalURL string `json:"ical_url"`
}
type Alert struct {
	Id     *uuid.UUID `json:"id"`      // identifiant unique de l'alerte
	Agenda int        `json:"agenda"`  // id de l'agenda associé (clé étrangère)
	Dest   string     `json:"dest"`    // destinataire
	Msg    string     `json:"msg"`     // message de l'alerte
	TypeA  string     `json:"typeA"`   // type d'alerte
}

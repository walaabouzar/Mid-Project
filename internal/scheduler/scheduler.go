package scheduler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Agenda minimal pour le scheduler
type Agenda struct {
	ID      int    `json:"id"`
	GroupID string `json:"group_id"`
	ICalURL string `json:"ical_url"`
}

// FetchAgendas récupère tous les agendas depuis l'API Config
func FetchAgendas(apiURL string) ([]Agenda, error) {
	resp, err := http.Get(apiURL + "/agendas")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var agendas []Agenda
	if err := json.Unmarshal(body, &agendas); err != nil {
		return nil, err
	}

	return agendas, nil
}

// AfficheAgendas est une fonction utilitaire pour logging
func AfficheAgendas(agendas []Agenda) {
	log.Println("Agendas récupérés :")
	for _, a := range agendas {
		log.Printf("ID=%d, Group=%s, URL=%s\n", a.ID, a.GroupID, a.ICalURL)
	}
}

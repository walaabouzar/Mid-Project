package main

import (
	"encoding/json"
	"fmt"
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

// fetchAgendas récupère tous les agendas depuis l'API Config
func fetchAgendas(apiURL string) ([]Agenda, error) {
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

func main() {
	// Adresse de ton API Config
	apiConfigURL := "http://localhost:8080"

	agendas, err := fetchAgendas(apiConfigURL)
	if err != nil {
		log.Fatal("Erreur fetch agendas:", err)
	}

	fmt.Println("Agendas récupérés :")
	for _, a := range agendas {
		fmt.Printf("ID=%d, Group=%s, URL=%s\n", a.ID, a.GroupID, a.ICalURL)
	}
}

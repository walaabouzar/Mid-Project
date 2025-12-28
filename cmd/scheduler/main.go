package main

import (
	"log"
	"middleware/example/internal/scheduler"
)

func main() {
	apiConfigURL := "http://localhost:8080" // URL de ton API Config

	// Récupération des agendas
	agendas, err := scheduler.FetchAgendas(apiConfigURL)
	if err != nil {
		log.Fatal("Erreur fetch agendas:", err)
	}

	// Affichage des agendas récupérés
	scheduler.AfficheAgendas(agendas)

	// Pour chaque agenda, récupérer les événements iCal
	for _, a := range agendas {
		events, err := scheduler.FetchICalEvents(a.ICalURL)
		if err != nil {
			log.Println("Erreur fetch iCal pour agenda", a.GroupID, ":", err)
			continue
		}

		log.Printf("Événements pour %s (%d) :", a.GroupID, a.ID)
		scheduler.AfficheEvents(events)
	}
}

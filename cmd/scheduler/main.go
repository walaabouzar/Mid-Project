package main

import (
	"log"
	"time"

	"middleware/example/internal/scheduler" // Remplace par le chemin réel de ton module
)

func main() {
	apiURL := "http://localhost:8080" // Remplace par l’URL de ton API Config
	interval := 5 * time.Second       // Intervalle de 5 secondes

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Première exécution immédiate
	runScheduler(apiURL)

	for range ticker.C {
		runScheduler(apiURL)
	}
}

func runScheduler(apiURL string) {
	log.Println("=== Début du Scheduler ===")

	// 1. Récupérer les agendas
	agendas, err := scheduler.FetchAgendas(apiURL)
	if err != nil {
		log.Println("Erreur fetch agendas:", err)
		return
	}
	scheduler.AfficheAgendas(agendas)

	// 2. Pour chaque agenda, récupérer les events et les afficher
	for _, a := range agendas {
		log.Printf("Événements pour %s (%d) :", a.GroupID, a.ID)
		events, err := scheduler.FetchICalEvents(a.ICalURL)
		if err != nil {
			log.Println("Erreur fetch iCal:", err)
			continue
		}
		scheduler.AfficheEvents(events)
	}

	log.Println("=== Fin du Scheduler ===")
}

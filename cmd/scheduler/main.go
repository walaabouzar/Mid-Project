package main

import (
	"log"
	"time"

	"middleware/example/internal/scheduler"
)

func main() {
	scheduler.InitNats()

	apiURL := "http://localhost:8080"
	interval := 5 * time.Second

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Lancer la première exécution immédiatement
	runScheduler(apiURL)

	for range ticker.C {
		runScheduler(apiURL)
	}
}

func runScheduler(apiURL string) {
	log.Println("=== Début du Scheduler ===")

	agendas, err := scheduler.FetchAgendas(apiURL)
	if err != nil {
		log.Println("Erreur fetch agendas:", err)
		return
	}

	for _, a := range agendas {
		events, err := scheduler.FetchICalEvents(a.ICalURL, a.ID, a.GroupID)
		if err != nil {
			log.Println("Erreur fetch iCal:", err)
			continue
		}

		for _, e := range events {
			scheduler.PublishEvent(a.GroupID, e)
		}

		log.Printf("Publié %d événements pour le groupe %s sur NATS", len(events), a.GroupID)
	}

	log.Println("=== Fin du Scheduler ===")
}

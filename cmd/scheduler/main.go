package main

import (
	"log"
	"time"

	"middleware/example/internal/scheduler"
)

func main() {
	log.Println("Scheduler démarré")

	// NATS CENTRAL
	scheduler.InitNats("nats://localhost:4222")

	API_URL := "http://localhost:8080"

	for {
		agendas, err := scheduler.FetchAgendas(API_URL)
		if err != nil {
			log.Println("Erreur fetch agendas:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		for _, agenda := range agendas {
			events, err := scheduler.FetchICalEvents(
				agenda.ICalURL,
				agenda.ID,
				agenda.GroupID,
			)
			if err != nil {
				log.Println("Erreur iCal:", err)
				continue
			}

			for _, event := range events {
				scheduler.PublishEvent(agenda.GroupID, event)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

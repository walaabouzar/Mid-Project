package main

import (
	"log"
	"middleware/example/internal/scheduler"
)

func main() {
	apiConfigURL := "http://localhost:8080" // URL de API Config

	agendas, err := scheduler.FetchAgendas(apiConfigURL)
	if err != nil {
		log.Fatal("Erreur fetch agendas:", err)
	}

	scheduler.AfficheAgendas(agendas)

	
}

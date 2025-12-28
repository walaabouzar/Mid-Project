package main

import (
	"fmt"
)

// Agenda minimal pour le scheduler
type Agenda struct {
	ID      int
	GroupID string
	ICalURL string
}

// fetchAgendas simule la récupération des agendas depuis l'API Config
func fetchAgendas() []Agenda {
	// On simule un agenda comme celui de ton API Config
	return []Agenda{
		{ID: 31, GroupID: "G1", ICalURL: "https://example.com/calendar1.ics"},
		{ID: 32, GroupID: "G2", ICalURL: "https://example.com/calendar2.ics"},
	}
}

func main() {
	agendas := fetchAgendas()

	fmt.Println("Agendas récupérés :")
	for _, a := range agendas {
		fmt.Printf("ID=%d, Group=%s, URL=%s\n", a.ID, a.GroupID, a.ICalURL)
	}
}

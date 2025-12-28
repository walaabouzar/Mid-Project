package scheduler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	ical "github.com/arran4/golang-ical"
)

// --- Agendas ---

func FetchAgendas(apiURL string) ([]Agenda, error) {
	resp, err := http.Get(apiURL + "/agendas")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var agendas []Agenda
	if err := json.Unmarshal(data, &agendas); err != nil {
		return nil, err
	}

	return agendas, nil
}

func AfficheAgendas(agendas []Agenda) {
	log.Println("Agendas récupérés :")
	for _, a := range agendas {
		log.Printf("ID=%d, Group=%s, URL=%s\n", a.ID, a.GroupID, a.ICalURL)
	}
}

// --- iCal ---

func FetchICalEvents(url string) ([]Event, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cal, err := ical.ParseCalendar(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var events []Event
	for _, e := range cal.Events() {
		start, err := e.GetStartAt()
		if err != nil {
			log.Println("Erreur parse start:", err)
			continue
		}

		end, err := e.GetEndAt()
		if err != nil {
			log.Println("Erreur parse end:", err)
			continue
		}

		summary := ""
		if prop := e.GetProperty(ical.ComponentPropertySummary); prop != nil {
			summary = prop.Value
		}

		events = append(events, Event{
			Summary: summary,
			Start:   start,
			End:     end,
		})
	}

	return events, nil
}

func AfficheEvents(events []Event) {
	for _, e := range events {
		log.Printf(
			"%s : %s → %s\n",
			e.Summary,
			e.Start.Format(time.RFC3339),
			e.End.Format(time.RFC3339),
		)
	}
}

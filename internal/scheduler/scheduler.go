package scheduler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	ical "github.com/arran4/golang-ical"
	"github.com/nats-io/nats.go"
)

/* =======================
   NATS
======================= */

var js nats.JetStreamContext

func InitNats(natsURL string) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Erreur connexion NATS:", err)
	}

	js, err = nc.JetStream()
	if err != nil {
		log.Fatal("Erreur JetStream:", err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "COURS",
		Subjects: []string{"COURS.>"},
	})
	if err != nil {
		log.Println("Stream COURS déjà existant")
	}
}

func PublishEvent(groupID string, event Event) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Erreur marshal event:", err)
		return
	}

	subject := "COURS." + strings.ReplaceAll(groupID, " ", ".")
	_, err = js.Publish(subject, data)
	if err != nil {
		log.Println("Erreur publish NATS:", err)
	}
}

/* =======================
   AGENDAS
======================= */

func FetchAgendas(apiURL string) ([]Agenda, error) {
	resp, err := http.Get(apiURL + "/agendas")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var agendas []Agenda
	err = json.Unmarshal(body, &agendas)
	return agendas, err
}

/* =======================
   ICAL
======================= */

func FetchICalEvents(url string, agendaID int, groupID string) ([]Event, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	cal, err := ical.ParseCalendar(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var events []Event
	for _, e := range cal.Events() {
		start, _ := e.GetStartAt()
		end, _ := e.GetEndAt()

		summary := ""
		if p := e.GetProperty(ical.ComponentPropertySummary); p != nil {
			summary = p.Value
		}

		events = append(events, Event{
			AgendaID: agendaID,
			GroupID:  groupID,
			Summary:  summary,
			Start:    start,
			End:      end,
		})
	}

	return events, nil
}

/* =======================
   DEBUG
======================= */

func PrintEvents(events []Event) {
	for _, e := range events {
		log.Printf(
			"%s | %s → %s | agenda=%d | group=%s",
			e.Summary,
			e.Start.Format(time.RFC3339),
			e.End.Format(time.RFC3339),
			e.AgendaID,
			e.GroupID,
		)
	}
}

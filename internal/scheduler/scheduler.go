package scheduler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings" // Ajouté pour le nettoyage des sujets
	"time"

	ical "github.com/arran4/golang-ical"
	"github.com/nats-io/nats.go"
)

// Variable globale pour JetStream
var js nats.JetStreamContext

// --- NATS Functions ---

func InitNats() {
	nc, err := nats.Connect(nats.DefaultURL)
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
		log.Println("Note: Le stream existe déjà ou erreur lors de la création.")
	}
}

func PublishEvent(groupID string, event Event) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Erreur marshal event:", err)
		return
	}

	// NETTOYAGE : NATS n'accepte pas les espaces dans les sujets.
	// On remplace les espaces par des points pour créer une hiérarchie valide.
	// Exemple: "M1 G1 Langue" -> "M1.G1.Langue"
	cleanGroupID := strings.ReplaceAll(groupID, " ", ".")

	subject := "COURS." + cleanGroupID
	_, err = js.Publish(subject, data)
	if err != nil {
		log.Printf("Erreur publication NATS sur %s: %v\n", subject, err)
	}
}

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
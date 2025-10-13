package event

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"middleware/example/internal/models"

	"github.com/emersion/go-ical"
)
// Liste des agendas (tu peux les mettre dans un fichier config si tu veux)
var agendas = []models.Agenda{
	{ID: 13295, Name: "M1 Groupe 1 langue"},
	{ID: 13345, Name: "M1 Groupe 2 langue"},
	{ID: 13397, Name: "M1 Groupe 3 langue"},
	{ID: 7224, Name: "M1 Groupe 1 option"},
	{ID: 7225, Name: "M1 Groupe 2 option"},
	{ID: 62962, Name: "M1 Groupe 3 option"},
	{ID: 62090, Name: "M1 Groupe option"},
	{ID: 56529, Name: "M1 - Tutorat L2"},
}

// URL de base du calendrier
const baseURL = "https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%d&projectId=3&calType=ical&nbWeeks=8&displayConfigId=128"

// FetchAndParseEvents télécharge et parse tous les événements ICS
func FetchAndParseEvents(w http.ResponseWriter, r *http.Request) {
	var allEvents []models.Event

	for _, agenda := range agendas {
		url := fmt.Sprintf(baseURL, agenda.ID)
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Erreur de téléchargement: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Erreur HTTP "+resp.Status, http.StatusInternalServerError)
			return
		}

		events, err := parseICS(resp.Body, agenda.ID)
		if err != nil {
			http.Error(w, "Erreur parsing ICS: "+err.Error(), http.StatusInternalServerError)
			return
		}

		allEvents = append(allEvents, events...)
	}

	// Conversion en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allEvents)
}

// parseICS lit et transforme le contenu ICS en slice d'Event
// parseICS lit et transforme le contenu ICS en slice d'Event
func parseICS(body io.Reader, agendaID int) ([]models.Event, error) {
	var events []models.Event
	
	cal, err := ical.NewDecoder(body).Decode()
	if err != nil {
		return nil, err
	}
	
	for _, comp := range cal.Children {
		if comp.Name == ical.CompEvent {
			event := models.Event{AgendaID: agendaID}
			
			// Accès aux propriétés via Props au lieu de Children
			if prop := comp.Props.Get(ical.PropUID); prop != nil {
				event.UID = prop.Value
			}
			
			if prop := comp.Props.Get(ical.PropSummary); prop != nil {
				event.Title = prop.Value
			}
			
			if prop := comp.Props.Get(ical.PropLocation); prop != nil {
				event.Location = prop.Value
			}
			
			if prop := comp.Props.Get(ical.PropDescription); prop != nil {
				// Nettoyage des sauts de ligne \n
				event.Description = strings.ReplaceAll(prop.Value, "\\n", " ")
			}
			
			if prop := comp.Props.Get(ical.PropDateTimeStart); prop != nil {
				t, _ := time.Parse("20060102T150405Z", prop.Value)
				event.Start = t
			}
			
			if prop := comp.Props.Get(ical.PropDateTimeEnd); prop != nil {
				t, _ := time.Parse("20060102T150405Z", prop.Value)
				event.End = t
			}
			
			events = append(events, event)
		}
	}
	
	return events, nil
}
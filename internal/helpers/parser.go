package helpers

import (
	"strings"
	"time"

	"github.com/emersion/go-ical"
	"github.com/sirupsen/logrus"
	
	"middleware/example/internal/models"
)

// ParseEvents parse le contenu iCal et retourne une liste d'événements
func ParseEvents(icalContent string, agendaID string) ([]models.Event, error) {
	decoder := ical.NewDecoder(strings.NewReader(icalContent))
	
	var events []models.Event
	
	for {
		cal, err := decoder.Decode()
		if err != nil {
			break
		}
		
		for _, comp := range cal.Children {
			if comp.Name != "VEVENT" {
				continue
			}
			
			event := models.Event{
				AgendaID: agendaID,
			}
			
			// Parser les propriétés de l'événement
			// comp.Props est une map[string][]ical.Prop
			if uid := comp.Props.Get("UID"); uid != nil {
				event.UID = uid.Value
			}
			
			if summary := comp.Props.Get("SUMMARY"); summary != nil {
				event.Summary = summary.Value
			}
			
			if location := comp.Props.Get("LOCATION"); location != nil {
				event.Location = location.Value
			}
			
			if description := comp.Props.Get("DESCRIPTION"); description != nil {
				event.Description = description.Value
			}
			
			if dtstart := comp.Props.Get("DTSTART"); dtstart != nil {
				if t, err := parseICalTime(dtstart.Value); err == nil {
					event.StartTime = t
				}
			}
			
			if dtend := comp.Props.Get("DTEND"); dtend != nil {
				if t, err := parseICalTime(dtend.Value); err == nil {
					event.EndTime = t
				}
			}
			
			if lastMod := comp.Props.Get("LAST-MODIFIED"); lastMod != nil {
				if t, err := parseICalTime(lastMod.Value); err == nil {
					event.LastModified = t
				}
			}
			
			if seq := comp.Props.Get("SEQUENCE"); seq != nil {
				// Vous pouvez parser la séquence si nécessaire
				// event.Sequence = convertToInt(seq.Value)
			}
			
			events = append(events, event)
		}
	}
	
	logrus.Infof("✅ %d événements parsés pour AgendaID=%s", len(events), agendaID)
	return events, nil
}

// parseICalTime parse les différents formats de date/heure iCal
func parseICalTime(value string) (time.Time, error) {
	// Format avec Z (UTC)
	if t, err := time.Parse("20060102T150405Z", value); err == nil {
		return t, nil
	}
	
	// Format sans Z (local)
	if t, err := time.Parse("20060102T150405", value); err == nil {
		return t, nil
	}
	
	// Format date seule
	if t, err := time.Parse("20060102", value); err == nil {
		return t, nil
	}
	
	return time.Time{}, nil
}
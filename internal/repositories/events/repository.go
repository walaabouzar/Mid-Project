package events

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllEvents() ([]models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, agenda_id, uid, description, name, start, end, location, last_update FROM events")
	if err != nil {
		helpers.CloseDB(db)
		return nil, err
	}

	events := []models.Event{}
	for rows.Next() {
		var e models.Event
		err = rows.Scan(&e.ID, &e.AgendaID, &e.UID, &e.Description, &e.Name, &e.Start, &e.End, &e.Location, &e.LastUpdate)
		if err != nil {
			helpers.CloseDB(db)
			return nil, err
		}
		events = append(events, e)
	}
	_ = rows.Close()
	helpers.CloseDB(db)

	return events, nil
}

func GetEventById(id uuid.UUID) (*models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT id, agenda_id, uid, description, name, start, end, location, last_update FROM events WHERE id=?", id.String())

	var e models.Event
	err = row.Scan(&e.ID, &e.AgendaID, &e.UID, &e.Description, &e.Name, &e.Start, &e.End, &e.Location, &e.LastUpdate)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

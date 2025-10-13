package event

import (
	
	"log"
	

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func InsertEvent(e models.Event) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	if e.ID == nil {
		id, _ := uuid.NewV4()
		e.ID = &id
	}

	query := `
		INSERT INTO events (id, u_uid, title, location, description, start, end, agenda_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err = db.Exec(query, e.ID.String(), e.UID, e.Title, e.Location, e.Description, e.Start, e.End, e.AgendaID)
	if err != nil {
		return err
	}

	log.Println("✅ Événement inséré avec succès :", e.Title)
	return nil
}

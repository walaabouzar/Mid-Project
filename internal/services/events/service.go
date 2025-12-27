package events

import (
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/events"
)

func GetAllEvents() ([]models.Event, error) {
	events, err := repository.GetAllEvents()
	if err != nil {
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving events",
		}
	}
	return events, nil
}

func GetEventById(id uuid.UUID) (*models.Event, error) {
	event, err := repository.GetEventById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorNotFound{
				Message: fmt.Sprintf("Event %s not found", id.String()),
			}
		}
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving event %s", id.String()),
		}
	}
	return event, nil
}

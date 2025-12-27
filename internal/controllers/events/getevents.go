package events

import (
	"encoding/json"
	"net/http"

	"middleware/example/internal/helpers"
	"middleware/example/internal/services/events"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	allEvents, err := events.GetAllEvents()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(allEvents)
	_, _ = w.Write(body)
}

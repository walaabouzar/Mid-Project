package alerts

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/alerts"
)
// CreateAgenda godoc
// @Summary Create a new agenda
// @Description Add a new agenda to the database
// @Tags agendas
// @Accept json
// @Produce json
// @Param agenda body models.Agenda true "Agenda info"
// @Success 201 {object} models.Agenda
// @Failure 400 {object} map[string]string
// @Router /agendas [post]
// CreateAgendaHandler crée un nouvel agenda
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var newAlert models.Alert

	// Décoder le body JSON
	err := json.NewDecoder(r.Body).Decode(&newAlert)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// Appel au service pour créer l'agenda
	createdAlert, err := alerts.CreateAlert(newAlert)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// Réponse OK
	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(createdAlert)
	_, _ = w.Write(body)
}

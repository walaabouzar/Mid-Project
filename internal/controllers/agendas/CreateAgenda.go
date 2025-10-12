package agendas

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/agendas"
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
func CreateAgenda(w http.ResponseWriter, r *http.Request) {
	var newAgenda models.Agenda

	// Décoder le body JSON
	err := json.NewDecoder(r.Body).Decode(&newAgenda)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// Appel au service pour créer l'agenda
	createdAgenda, err := agendas.CreateAgenda(newAgenda)
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
	body, _ := json.Marshal(createdAgenda)
	_, _ = w.Write(body)
}

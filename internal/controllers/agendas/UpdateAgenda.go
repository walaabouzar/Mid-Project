package agendas

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/agendas"
)

// UpdateAgendaHandler met à jour un agenda existant
func UpdateAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaID, _ := ctx.Value("agendaID").(int)

	var agenda models.Agenda
	err := json.NewDecoder(r.Body).Decode(&agenda)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	updatedAgenda, err := agendas.UpdateAgenda(agendaID, agenda)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	body, _ := json.Marshal(updatedAgenda)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

package agendas

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/repositories/agendas"
)
// GetAgenda godoc
// @Summary Get an agenda by ID
// @Description Retrieve a single agenda by ID
// @Tags agendas
// @Param id path int true "Agenda ID"
// @Success 200 {object} models.Agenda
// @Failure 404 {object} map[string]string
// @Router /agendas/{id} [get]

func GetAgendaByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaID, _ := ctx.Value("agendaID").(int)

	agenda, err := agendas.GetAgendaByID(agendaID)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(agenda)
	_, _ = w.Write(body)
}

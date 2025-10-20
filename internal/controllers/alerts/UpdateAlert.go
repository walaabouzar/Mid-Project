package alerts

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/alerts"
	"github.com/gofrs/uuid"
)

// UpdateAgenda godoc
// @Summary Update an agenda
// @Description Update an agenda by ID
// @Tags agendas
// @Accept json
// @Produce json
// @Param id path int true "Agenda ID"
// @Param agenda body models.Agenda true "Updated agenda info"
// @Success 200 {object} models.Agenda
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /agendas/{id} [put]


// UpdateAgendaHandler met à jour un agenda existant
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertID, _ := ctx.Value("alertId").(uuid.UUID)

	var alert models.Alert
	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	updatedAlert, err := alerts.UpdateAlert(alertID, alert)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	body, _ := json.Marshal(updatedAlert)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

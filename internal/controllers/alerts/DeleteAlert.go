package alerts

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/alerts"
	"github.com/gofrs/uuid"	
)

// DeleteAgenda godoc
// @Summary Delete an agenda
// @Description Delete an agenda by ID
// @Tags agendas
// @Param id path int true "Agenda ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /agendas/{id} [delete]

// DeleteAgendaHandler supprime un agenda par son ID
func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertID, _ := ctx.Value("alertId").(uuid.UUID)

	err := alerts.DeleteAlert(alertID)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	response := map[string]string{"message": "Alert deleted successfully"}
	body, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

package alerts

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/alerts"
)

// DeleteAllAgendas godoc
// @Summary Delete all agendas
// @Description Delete all agendas from the database
// @Tags agendas
// @Success 200 {object} map[string]string
// @Router /agendas [delete]
// DeleteAllAgendasHandler supprime tous les agendas de la base
func DeleteAllAlerts(w http.ResponseWriter, r *http.Request) {
	err := alerts.DeleteAllAlerts()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	response := map[string]string{"message": "All alerts deleted successfully"}
	body, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

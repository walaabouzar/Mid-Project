package alerts

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/alerts"
	"net/http"
)

// GetUsers
// @Tags         users
// @Summary      Get all users.
// @Description  Get all users.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [get]
func GetAllAlerts(w http.ResponseWriter, _ *http.Request) {
	// calling service
	alerts, err := alerts.GetAllAlerts()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alerts)
	_, _ = w.Write(body)
	return
}

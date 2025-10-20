package alerts

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/alerts"
	"net/http"
)

// GetUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [get]
func GetAlertById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	AlertId, _ := ctx.Value("alertId").(uuid.UUID) // getting key set in context.go

	alert, err := alerts.GetAlertById(AlertId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alert)
	_, _ = w.Write(body)
	return
}

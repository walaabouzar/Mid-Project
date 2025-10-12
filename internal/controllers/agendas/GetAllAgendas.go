package agendas

import (
	"encoding/json"
	"net/http"
	"middleware/example/internal/helpers"
	"middleware/example/internal/repositories/agendas"
)

// GetAllAgendas récupère tous les agendas
func GetAllAgendas(w http.ResponseWriter, r *http.Request) {
	allAgendas, err := agendas.GetAllAgendas()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(allAgendas)
	_, _ = w.Write(body)
}

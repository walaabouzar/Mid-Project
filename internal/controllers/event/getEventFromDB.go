package event

import (
	"encoding/json"
	"net/http"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"github.com/sirupsen/logrus"
)

// GetEventsFromDB : récupère tous les événements déjà stockés dans la base SQLite
func GetEventsFromDB(w http.ResponseWriter, r *http.Request) {
	db, err := helpers.OpenDB()
	if err != nil {
		http.Error(w, "Erreur de connexion à la base de données", http.StatusInternalServerError)
		logrus.Error("Erreur connexion DB :", err)
		return
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query(`SELECT id, u_uid, title, location, description, start, end, agenda_id FROM events`)
	if err != nil {
		http.Error(w, "Erreur lors de la requête SQL", http.StatusInternalServerError)
		logrus.Error("Erreur requête SELECT :", err)
		return
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.ID, &e.UID, &e.Title, &e.Location, &e.Description, &e.Start, &e.End, &e.AgendaID); err != nil {
			logrus.Error("Erreur lecture ligne :", err)
			continue
		}
		events = append(events, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

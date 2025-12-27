package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid" // <-- Ajout ici
	"github.com/sirupsen/logrus"
	"net/http"

	"middleware/example/internal/controllers/events"
	"middleware/example/internal/helpers"
)

func main() {
	r := chi.NewRouter()

	r.Route("/events", func(r chi.Router) {
		r.Get("/", events.GetEvents) // GET /events

		r.Route("/{id}", func(r chi.Router) {
			r.Use(events.Context)
			r.Get("/", events.GetEvent) // GET /events/{id}
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8081", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	scheme := `CREATE TABLE IF NOT EXISTS events (
		id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
		agenda_id INTEGER,
		uid VARCHAR(255) NOT NULL,
		description TEXT,
		name VARCHAR(255) NOT NULL,
		start DATETIME,
		end DATETIME,
		location VARCHAR(255),
		last_update DATETIME
	);`
	if _, err := db.Exec(scheme); err != nil {
		logrus.Fatalln("Could not generate events table ! Error was : " + err.Error())
	}

	// Insertion d'un seul événement test avec UUID généré automatiquement
	id := uuid.Must(uuid.NewV4()).String() // Génération UUID
	_, _ = db.Exec(`INSERT OR IGNORE INTO events 
		(id, agenda_id, uid, description, name, start, end, location, last_update)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		id,
		13295,
		"ADE60323032352d323032362d5543412d32353835322d312d30",
		"*TD Culture d'entreprise G2\nM1 GROUPE 2 langue\nLACHENAUD SOPHIE\n(Updated :09/10/2025 18:32)",
		"TD Culture d'entreprise G2",
		"2025-12-03T10:00:00+00:00",
		"2025-12-03T11:30:00+00:00",
		"IS_A002",
		"2025-10-09T16:32:00+00:00",
	)

	helpers.CloseDB(db)
}

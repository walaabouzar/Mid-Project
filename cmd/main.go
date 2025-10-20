package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/controllers/agendas"
	"middleware/example/internal/controllers/alerts"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// ---------- USERS ----------
	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Context)
			r.Get("/", users.GetUser)
		})
	})

	// ---------- AGENDAS ----------
	r.Route("/agendas", func(r chi.Router) {
		r.Get("/", agendas.GetAllAgendas)     // GET /agendas
		r.Post("/", agendas.CreateAgenda)     // POST /agendas
		r.Delete("/", agendas.DeleteAllAgendas) // DELETE /agendas

		r.Route("/{id}", func(r chi.Router) {
			r.Use(agendas.ContextAgenda)
			r.Get("/", agendas.GetAgendaByID) // GET /agendas/{id}
			r.Put("/", agendas.UpdateAgenda)  // PUT /agendas/{id}
			r.Delete("/", agendas.DeleteAgenda) // DELETE /agendas/{id}
		})
	})
	// ---------- ALERTS ----------
	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", alerts.GetAllAlerts)       // GET /alerts
		r.Post("/", alerts.CreateAlert)       // POST /alerts
		r.Delete("/", alerts.DeleteAllAlerts) // DELETE /alerts

		r.Route("/{id}", func(r chi.Router) {
			r.Use(alerts.ContextAlerts)        // middleware pour récupérer l'ID de l'alerte
			r.Get("/", alerts.GetAlertById)   // GET /alerts/{id}
			r.Put("/", alerts.UpdateAlert)    // PUT /alerts/{id}
			r.Delete("/", alerts.DeleteAlert) // DELETE /alerts/{id}
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB() // ✅ sans argument
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS agendas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			group_id TEXT NOT NULL,
			ical_url TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS alerts (
			Id VARCHAR(255) PRIMARY KEY,                 -- Identifiant unique de l’alerte
			Dest VARCHAR(255) NOT NULL,                   -- Adresse e-mail de l'étudiant
			Msg TEXT NOT NULL,                        -- Contenu du message de l’alerte
			TypeA VARCHAR(100) NOT NULL,                   -- Type de l’alerte (ex : reminder, deadline, update)
			Agenda INTEGER NOT NULL REFERENCES agendas(id) ON DELETE CASCADE -- Clé étrangère vers un agenda
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error: " + err.Error())
		}
	}

	helpers.CloseDB(db)
}

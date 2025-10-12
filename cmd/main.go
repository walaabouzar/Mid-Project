package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/controllers/agendas"
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
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error: " + err.Error())
		}
	}

	helpers.CloseDB(db)
}

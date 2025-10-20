package main

import (
	"net/http"
	//"time"
	  
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	"middleware/example/internal/controllers/event"
	//"middleware/example/internal/models"
)

func main() {
	db, _ := helpers.OpenDB()
	defer helpers.CloseDB(db)

/*	testEvent := models.Event{
		UID:         "ADE60323032352d323032362d5543412d35383437302d302d30",
		Title:       "Examen Calculabilité (1/3 temps)",
		Location:    "Salle 101",
		Description: "MASTER 1 INFO RAYNAUD OLIVIER",
		Start:       time.Now(),
		End:         time.Now().Add(2 * time.Hour),
		AgendaID:    13295,
	}

	if err := event.InsertEvent(testEvent); err != nil {
		logrus.Fatalf("Erreur insertion : %s", err.Error())
	}

	logrus.Info("✅ Test d’insertion réussi !")
	
*/
	
	r := chi.NewRouter()
	
	r.Route("/users", func(r chi.Router) { // route /users
		r.Get("/", users.GetUsers)            // GET /users
		r.Route("/{id}", func(r chi.Router) { // route /users/{id}
			r.Use(users.Context)      // Use Context method to get user ID
			r.Get("/", users.GetUser) // GET /users/{id}
		})
	})
	
	
	// Route pour afficher tous les événements de la base
r.Get("/events/db", event.GetEventsFromDB)

	// Route pour récupérer et parser tous les événements
	r.Get("/events", event.FetchAndParseEvents)
	
	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	
	// AJOUTER CETTE LIGNE pour démarrer le serveur
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Fatalf("Error starting server: %s", err.Error())
	}


}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS events (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			u_uid VARCHAR(255) NOT NULL,
			title VARCHAR(255),
			location VARCHAR(255),
			description TEXT,
			start DATETIME,
			end DATETIME,
			agenda_id INTEGER
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error was: " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
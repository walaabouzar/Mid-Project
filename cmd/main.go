/*package main
*/
/*
import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)
*/
/*
func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) { // route /users
		r.Get("/", users.GetUsers)            // GET /users
		r.Route("/{id}", func(r chi.Router) { // route /users/{id}
			r.Use(users.Context)      // Use Context method to get user ID
			r.Get("/", users.GetUser) // GET /users/{id}
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
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
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
	
}
*/
package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// Routes utilisateurs existantes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Context)
			r.Get("/", users.GetUser)
		})
	})
// Ajoutez cette route dans votre main.go pour analyser les doublons

r.Get("/check-duplicates", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	db, err := helpers.OpenDB()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("❌ Erreur connexion DB : %v\n", err)))
		return
	}
	defer helpers.CloseDB(db)
	
	var output string
	output += "===============================================\n"
	output += "   ANALYSE DES ÉVÉNEMENTS\n"
	output += "===============================================\n\n"
	
	// Vérifier s'il y a des UIDs qui apparaissent dans plusieurs agendas
	rows, err := db.Query(`
		SELECT id, COUNT(DISTINCT agenda_id) as agenda_count, 
		       GROUP_CONCAT(DISTINCT agenda_id) as agendas
		FROM events
		GROUP BY id
		HAVING COUNT(DISTINCT agenda_id) > 1
		ORDER BY agenda_count DESC
	`)
	
	if err != nil {
		output += fmt.Sprintf("❌ Erreur requête : %v\n", err)
	} else {
		defer rows.Close()
		duplicateCount := 0
		
		output += "🔍 Événements présents dans plusieurs agendas :\n\n"
		
		for rows.Next() {
			var uid, agendas string
			var count int
			if err := rows.Scan(&uid, &count, &agendas); err != nil {
				continue
			}
			duplicateCount++
			if duplicateCount <= 10 {
				output += fmt.Sprintf("  • UID: %s\n", uid)
				output += fmt.Sprintf("    Dans %d agendas : %s\n\n", count, agendas)
			}
		}
		
		if duplicateCount == 0 {
			output += "  ✅ Aucun événement dupliqué trouvé\n\n"
		} else {
			output += fmt.Sprintf("⚠️  Total: %d événements sont partagés entre plusieurs agendas\n\n", duplicateCount)
		}
	}
	
	// Statistiques détaillées
	output += "--- Statistiques détaillées ---\n\n"
	
	statsRows, err := db.Query(`
		SELECT 
			a.name,
			a.id,
			COUNT(e.id) as event_count
		FROM agendas a
		LEFT JOIN events e ON a.id = e.agenda_id
		GROUP BY a.id, a.name
		ORDER BY event_count DESC
	`)
	
	if err != nil {
		output += fmt.Sprintf("❌ Erreur stats : %v\n", err)
	} else {
		defer statsRows.Close()
		totalExpected := 0
		totalUnique := 0
		
		for statsRows.Next() {
			var name, id string
			var count int
			if err := statsRows.Scan(&name, &id, &count); err != nil {
				continue
			}
			output += fmt.Sprintf("  %s (ID: %s)\n", name, id)
			output += fmt.Sprintf("    → %d événements\n\n", count)
			totalExpected += count
		}
		
		// Compter les événements uniques
		db.QueryRow("SELECT COUNT(DISTINCT id) FROM events").Scan(&totalUnique)
		
		output += fmt.Sprintf("📊 Total événements par agenda : %d\n", totalExpected)
		output += fmt.Sprintf("📊 Total événements uniques (UID) : %d\n", totalUnique)
		
		if totalExpected != totalUnique {
			output += fmt.Sprintf("\n⚠️  Différence de %d événements\n", totalExpected-totalUnique)
			output += "   → Cela signifie que certains événements sont comptés\n"
			output += "      dans plusieurs agendas (événements communs)\n"
		}
	}
	
	output += "\n===============================================\n"
	w.Write([]byte(output))
})
// Ajoutez cette route dans votre main.go, dans la fonction main() après les autres routes

// Route pour vérifier le contenu de la base de données
// Ajoutez cette route dans votre main.go, dans la fonction main() après les autres routes

// Route pour vérifier le contenu de la base de données
r.Get("/db-stats", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	db, err := helpers.OpenDB()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("❌ Erreur connexion DB : %v\n", err)))
		return
	}
	defer helpers.CloseDB(db)
	
	var output string
	output += "===============================================\n"
	output += "   STATISTIQUES DE LA BASE DE DONNÉES\n"
	output += "===============================================\n\n"
	
	// Compter les agendas
	var agendaCount int
	err = db.QueryRow("SELECT COUNT(*) FROM agendas").Scan(&agendaCount)
	if err != nil {
		output += fmt.Sprintf("❌ Erreur comptage agendas : %v\n", err)
	} else {
		output += fmt.Sprintf("📅 Nombre d'agendas : %d\n", agendaCount)
	}
	
	// Compter les événements
	var eventCount int
	err = db.QueryRow("SELECT COUNT(*) FROM events").Scan(&eventCount)
	if err != nil {
		output += fmt.Sprintf("❌ Erreur comptage events : %v\n", err)
	} else {
		output += fmt.Sprintf("📆 Nombre d'événements : %d\n", eventCount)
	}
	
	output += "\n--- Détail par agenda ---\n\n"
	
	// Détail par agenda
	rows, err := db.Query(`
		SELECT a.id, a.name, COUNT(e.uid) as event_count
		FROM agendas a
		LEFT JOIN events e ON a.id = e.agenda_id
		GROUP BY a.id, a.name
		ORDER BY a.name
	`)
	if err != nil {
		output += fmt.Sprintf("❌ Erreur requête détails : %v\n", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, name string
			var count int
			if err := rows.Scan(&id, &name, &count); err != nil {
				continue
			}
			output += fmt.Sprintf("  • %s (ID: %s)\n", name, id)
			output += fmt.Sprintf("    → %d événements\n\n", count)
		}
	}
	
	// Quelques exemples d'événements
	output += "\n--- Exemples d'événements ---\n\n"
	eventRows, err := db.Query(`
		SELECT summary, location, start_time, agenda_id
		FROM events
		ORDER BY start_time
		LIMIT 5
	`)
	if err != nil {
		output += fmt.Sprintf("❌ Erreur requête events : %v\n", err)
	} else {
		defer eventRows.Close()
		i := 1
		for eventRows.Next() {
			var summary, location, startTime, agendaID string
			if err := eventRows.Scan(&summary, &location, &startTime, &agendaID); err != nil {
				continue
			}
			output += fmt.Sprintf("%d. %s\n", i, summary)
			output += fmt.Sprintf("   📍 %s\n", location)
			output += fmt.Sprintf("   🕐 %s\n", startTime)
			output += fmt.Sprintf("   📅 Agenda: %s\n\n", agendaID)
			i++
		}
	}
	
	output += "===============================================\n"
	w.Write([]byte(output))
})
	// Route pour tester tous les agendas
	r.Get("/test-agendas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		var output string

		output += "===============================================\n"
		output += "   TEST DE TOUS LES AGENDAS\n"
		output += "===============================================\n\n"

		totalAgendas := len(models.Agendas)
		successCount := 0
		failCount := 0
		totalEvents := 0

		for i, agenda := range models.Agendas {
			output += fmt.Sprintf("\n[%d/%d] Test de l'agenda : %s\n", i+1, totalAgendas, agenda.Name)
			output += fmt.Sprintf("    AgendaID : %s\n", agenda.AgendaID)
			output += "    ----------\n"
			output += "    Téléchargement... "

			// Télécharger le calendrier
			icalContent, err := helpers.FetchCalendar(agenda.AgendaID)
			if err != nil {
				output += fmt.Sprintf("❌\n       Erreur : %v\n", err)
				failCount++
				continue
			}
			output += fmt.Sprintf("✅ (%d octets)\n", len(icalContent))

			// Parser les événements
			output += "    Parsing... "
			events, err := helpers.ParseEvents(icalContent, agenda.AgendaID)
			if err != nil {
				output += fmt.Sprintf("❌\n       Erreur : %v\n", err)
				failCount++
				continue
			}
			output += fmt.Sprintf("✅ (%d événements)\n", len(events))

			// Sauvegarder agenda et événements
			if err := helpers.SaveAgenda(agenda); err != nil {
				logrus.Warnf("Erreur sauvegarde agenda %s : %v", agenda.Name, err)
			}
			if err := helpers.SaveEvents(events); err != nil {
				logrus.Warnf("Erreur sauvegarde events pour agenda %s : %v", agenda.Name, err)
			}

			successCount++
			totalEvents += len(events)

			// Aperçu des événements
			if len(events) > 0 {
				output += "    Aperçu des événements :\n"
				limit := 3
				if len(events) < limit {
					limit = len(events)
				}
				for j := 0; j < limit; j++ {
					event := events[j]
					output += fmt.Sprintf("       • %s\n", event.Summary)
					output += fmt.Sprintf("         %s\n", event.Location)
					output += fmt.Sprintf("         %s → %s\n",
						event.StartTime.Format("02/01 15:04"),
						event.EndTime.Format("15:04"))
				}
				if len(events) > limit {
					output += fmt.Sprintf("       ... et %d autres événements\n", len(events)-limit)
				}
			} else {
				output += "    ⚠️  Aucun événement trouvé (agenda vide ?)\n"
			}

			time.Sleep(200 * time.Millisecond)
		}

		// Résumé final
		output += "\n\n===============================================\n"
		output += "   RÉSUMÉ DU TEST\n"
		output += "===============================================\n"
		output += fmt.Sprintf("   Total agendas testés  : %d\n", totalAgendas)
		output += fmt.Sprintf("   ✅ Succès             : %d\n", successCount)
		output += fmt.Sprintf("   ❌ Échecs             : %d\n", failCount)
		output += fmt.Sprintf("   Total événements      : %d\n", totalEvents)

		if successCount == totalAgendas {
			output += "\n   🎉 TOUS LES TESTS SONT PASSÉS ! 🎉\n"
		} else if successCount > 0 {
			output += "\n   ⚠️  Certains tests ont échoué\n"
		} else {
			output += "\n   ❌ TOUS LES TESTS ONT ÉCHOUÉ\n"
		}

		output += "===============================================\n"
		w.Write([]byte(output))
	})

	logrus.Info("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	defer helpers.CloseDB(db)

	// ÉTAPE 1: Supprimer l'ancienne table events si elle existe
	logrus.Info("Migration: Suppression de l'ancienne table events...")
	_, err = db.Exec(`DROP TABLE IF EXISTS events;`)
	if err != nil {
		logrus.Warnf("Erreur lors de la suppression de la table events: %v", err)
	}

	// ÉTAPE 2: Créer les tables avec le bon schéma
	schemes := []string{
		// Table des utilisateurs
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL
		);`,

		// Table des agendas
		`CREATE TABLE IF NOT EXISTS agendas (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL
		);`,

		// Table des événements (avec clé composite)
	 `CREATE TABLE IF NOT EXISTS events (
    uid VARCHAR(255) NOT NULL,
    agenda_id VARCHAR(255) NOT NULL,
    summary TEXT,
    description TEXT,
    start_time DATETIME,
    end_time DATETIME,
    location TEXT,
    PRIMARY KEY (uid, agenda_id),
    FOREIGN KEY (agenda_id) REFERENCES agendas(id)
);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}

	logrus.Info("✅ Migration terminée: tables créées avec succès")
}

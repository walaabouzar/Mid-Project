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
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"
)

func main() {
	fmt.Println("離 ===============================================")
	fmt.Println("   TEST DE TOUS LES AGENDAS")
	fmt.Println("===============================================\n")
	
	// Compteurs globaux
	totalAgendas := len(models.Agendas)
	successCount := 0
	failCount := 0
	totalEvents := 0
	
	// Tester chaque agenda
	for i, agenda := range models.Agendas {
		fmt.Printf("\n [%d/%d] Test de l'agenda : %s\n", i+1, totalAgendas, agenda.Name)
		fmt.Printf("    AgendaID : %s\n", agenda.AgendaID)
		fmt.Println("    " + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]) + string([]rune("─")[0]))
		
		// Étape 1 : Télécharger
		fmt.Print("     Téléchargement... ")
		icalContent, err := helpers.FetchCalendar(agenda.AgendaID)
		if err != nil {
			fmt.Printf("❌\n       Erreur : %v\n", err)
			failCount++
			continue
		}
		fmt.Printf("✅ (%d octets)\n", len(icalContent))
		
		// Étape 2 : Parser
		fmt.Print("     Parsing... ")
		events, err := helpers.ParseEvents(icalContent, agenda.AgendaID)
		if err != nil {
			fmt.Printf("❌\n       Erreur : %v\n", err)
			failCount++
			continue
		}
		fmt.Printf("✅ (%d événements)\n", len(events))
		
		successCount++
		totalEvents += len(events)
		
		// Afficher quelques détails des événements
		if len(events) > 0 {
			fmt.Println("     Aperçu des événements :")
			limit := 3
			if len(events) < limit {
				limit = len(events)
			}
			
			for j := 0; j < limit; j++ {
				event := events[j]
				fmt.Printf("       • %s\n", event.Summary)
				fmt.Printf("          %s\n", event.Location)
				fmt.Printf("          %s → %s\n", 
					event.StartTime.Format("02/01 15:04"),
					event.EndTime.Format("15:04"))
			}
			
			if len(events) > limit {
				fmt.Printf("       ... et %d autres événements\n", len(events)-limit)
			}
		} else {
			fmt.Println("    ⚠️  Aucun événement trouvé (agenda vide ?)")
		}
		
		// Petite pause pour ne pas surcharger le serveur
		time.Sleep(500 * time.Millisecond)
	}
	
	// Résumé final
	fmt.Println("\n\n ===============================================")
	fmt.Println("   RÉSUMÉ DU TEST")
	fmt.Println("===============================================")
	fmt.Printf("   Total agendas testés  : %d\n", totalAgendas)
	fmt.Printf("   ✅ Succès             : %d\n", successCount)
	fmt.Printf("   ❌ Échecs             : %d\n", failCount)
	fmt.Printf("    Total événements   : %d\n", totalEvents)
	
	if successCount == totalAgendas {
		fmt.Println("\n    TOUS LES TESTS SONT PASSÉS ! ")
	} else if successCount > 0 {
		fmt.Println("\n   ⚠️  Certains tests ont échoué")
	} else {
		fmt.Println("\n   ❌ TOUS LES TESTS ONT ÉCHOUÉ")
	}
	
	fmt.Println("===============================================\n")
}
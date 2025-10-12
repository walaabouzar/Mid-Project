package helpers

import (
	"middleware/example/internal/models"
	"github.com/sirupsen/logrus"
)

// SaveAgenda insère un agenda dans la table agendas
func SaveAgenda(a models.Agenda) error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer CloseDB(db)

	// Utiliser AgendaID comme clé primaire (id) dans la table
	_, err = db.Exec(
		`INSERT OR IGNORE INTO agendas (id, name) VALUES (?, ?)`,
		a.AgendaID, a.Name,
	)
	if err != nil {
		logrus.Errorf("Erreur insertion agenda %s : %v", a.AgendaID, err)
		return err
	}
	return nil
}

// SaveEvents insère une liste d'événements dans la table events
func SaveEvents(events []models.Event) error {
	// Gérer le cas où il n'y a aucun événement
	if len(events) == 0 {
		logrus.Info("ℹ️  Aucun événement à insérer (agenda vide)")
		return nil
	}

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer CloseDB(db)

	successCount := 0
	errorCount := 0
	agendaID := events[0].AgendaID

	for _, e := range events {
		_, err := db.Exec(
			`INSERT OR REPLACE INTO events 
			(uid, agenda_id, summary, description, start_time, end_time, location)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
			e.UID, e.AgendaID, e.Summary, e.Description, e.StartTime, e.EndTime, e.Location,
		)
		if err != nil {
			logrus.Warnf("Erreur insertion event %s : %v", e.UID, err)
			errorCount++
		} else {
			successCount++
		}
	}
	
	logrus.Infof("✅ %d événements insérés avec succès (agenda %s)", successCount, agendaID)
	if errorCount > 0 {
		logrus.Warnf("⚠️  %d événements en erreur", errorCount)
	}
	
	return nil
}
package alerts

import (
	"fmt"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	
)

func GetAllAlerts() ([]models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM alerts")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	alerts := []models.Alert{}
	for rows.Next() {
		var data models.Alert
		err = rows.Scan(&data.Id, &data.Agenda, &data.Dest, &data.Msg, &data.TypeA)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return alerts, err
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM alerts WHERE Id=?", id.String())
	helpers.CloseDB(db)

	var alert models.Alert
	err = row.Scan(&alert.Id, &alert.Agenda, &alert.Dest, &alert.Msg, &alert.TypeA)
	if err != nil {
		return nil, err
	}
	return &alert, err
}
// Vérifie si un agenda existe
func CheckAgendaExists(agendaID int) (bool, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return false, err
	}
	defer helpers.CloseDB(db)

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM agendas WHERE Agenda= ?)", agendaID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
// CreateAgenda ajoute un nouvel agenda
func CreateAlert(a models.Alert) (*models.Alert, error) {
	// 1️⃣ Ouvrir la base
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// 4️⃣ Insérer l'alerte
	query := `INSERT INTO alerts (Id, Agenda, Dest, Msg, TypeA) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, a.Id, a.Agenda, a.Dest, a.Msg, a.TypeA)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de l'alerte : %v", err)
	}

	// 5️⃣ Retourner l’objet créé
	return &a, nil
}

// UpdateAlert met à jour une alerte dans la DB
func UpdateAlert(a models.Alert) (*models.Alert, error) {
	// 1️⃣ Ouvrir DB
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)
	//verfie si id de lalert existe
	// 2️⃣ Vérifier que l'alerte existe
	var alertExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM alerts WHERE Id = ?)", a.Id).Scan(&alertExists)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la vérification de l'alerte : %v", err)
	}
	if !alertExists {
		return nil, models.ErrorUnprocessableEntity{
			Message: fmt.Sprintf("Aucune alerte trouvée avec l'ID %v", a.Id),
		}
	}

	// 2️⃣ Vérifier que l'agenda existe
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM agendas WHERE id = ?)", a.Agenda).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la vérification de l'agenda : %v", err)
	}
	if !exists {
		return nil, models.ErrorUnprocessableEntity{
			Message: fmt.Sprintf("L'agenda %d n'existe pas", a.Agenda),
		}
	}

	// 3️⃣ Mettre à jour l'alerte
	query := `
		UPDATE alerts 
		SET Agenda = ?, Dest = ?, Msg = ?, TypeA = ?
		WHERE Id = ?
	`
	result, err := db.Exec(query, a.Agenda, a.Dest, a.Msg, a.TypeA, a.Id)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la mise à jour de l'alerte : %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, models.ErrorUnprocessableEntity{
			Message: fmt.Sprintf("Aucune alerte trouvée avec l'ID %v", a.Id),
		}
	}

	return &a, nil
}

// DeleteAgenda supprime un agenda par son ID
func DeleteAlert(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("DELETE FROM alerts WHERE Id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return helpers.ErrNoRows
	}
	return nil
}

// DeleteAllAgendas supprime tous les agendas
func DeleteAllAlerts() error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM alerts")
	if err != nil {
		return err
	}
	return nil
}

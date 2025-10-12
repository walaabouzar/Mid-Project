package agendas

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

// GetAllAgendas récupère tous les agendas
func GetAllAgendas() ([]models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query("SELECT id, group_id, ical_url FROM agendas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda
	for rows.Next() {
		var a models.Agenda
		if err := rows.Scan(&a.ID, &a.GroupID, &a.ICalURL); err != nil {
			return nil, err
		}
		agendas = append(agendas, a)
	}

	return agendas, nil
}

// GetAgendaByID récupère un agenda par son ID
func GetAgendaByID(id int) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT id, group_id, ical_url FROM agendas WHERE id = ?", id)
	var agenda models.Agenda
	if err := row.Scan(&agenda.ID, &agenda.GroupID, &agenda.ICalURL); err != nil {
		return nil, err
	}
	return &agenda, nil
}

// CreateAgenda ajoute un nouvel agenda
func CreateAgenda(a models.Agenda) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("INSERT INTO agendas (group_id, ical_url) VALUES (?, ?)", a.GroupID, a.ICalURL)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	a.ID = int(id)
	return &a, nil
}

// UpdateAgenda met à jour un agenda existant
func UpdateAgenda(id int, a models.Agenda) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("UPDATE agendas SET group_id = ?, ical_url = ? WHERE id = ?", a.GroupID, a.ICalURL, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, helpers.ErrNoRows
	}

	a.ID = id
	return &a, nil
}

// DeleteAgenda supprime un agenda par son ID
func DeleteAgenda(id int) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("DELETE FROM agendas WHERE id = ?", id)
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
func DeleteAllAgendas() error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM agendas")
	if err != nil {
		return err
	}
	return nil
}

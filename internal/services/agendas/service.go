package agendas

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/agendas"
)

// GetAllAgendas récupère toutes les agendas
func GetAllAgendas() ([]models.Agenda, error) {
	agendas, err := repository.GetAllAgendas()
	if err != nil {
		logrus.Errorf("error retrieving agendas: %s", err.Error())
		return nil, fmt.Errorf("something went wrong while retrieving agendas")
	}
	return agendas, nil
}

// GetAgendaByID récupère une agenda par son ID
func GetAgendaByID(id int) (*models.Agenda, error) {
	agenda, err := repository.GetAgendaByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("agenda not found")
		}
		logrus.Errorf("error retrieving agenda %d: %s", id, err.Error())
		return nil, fmt.Errorf("something went wrong while retrieving agenda %d", id)
	}
	return agenda, nil
}

// CreateAgenda crée une nouvelle agenda
func CreateAgenda(agenda models.Agenda) (*models.Agenda, error) {
	newAgenda, err := repository.CreateAgenda(agenda)
	if err != nil {
		logrus.Errorf("error creating agenda: %s", err.Error())
		return nil, fmt.Errorf("something went wrong while creating agenda")
	}
	return newAgenda, nil
}

// UpdateAgenda met à jour une agenda existante
func UpdateAgenda(id int, agenda models.Agenda) (*models.Agenda, error) {
	updatedAgenda, err := repository.UpdateAgenda(id, agenda)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("agenda not found")
		}
		logrus.Errorf("error updating agenda %d: %s", id, err.Error())
		return nil, fmt.Errorf("something went wrong while updating agenda %d", id)
	}
	return updatedAgenda, nil
}

// DeleteAgenda supprime une agenda par son ID
func DeleteAgenda(id int) error {
	err := repository.DeleteAgenda(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("agenda not found")
		}
		logrus.Errorf("error deleting agenda %d: %s", id, err.Error())
		return fmt.Errorf("something went wrong while deleting agenda %d", id)
	}
	return nil
}

// DeleteAllAgendas supprime toutes les agendas
func DeleteAllAgendas() error {
	err := repository.DeleteAllAgendas()
	if err != nil {
		logrus.Errorf("error deleting all agendas: %s", err.Error())
		return fmt.Errorf("something went wrong while deleting all agendas")
	}
	return nil
}

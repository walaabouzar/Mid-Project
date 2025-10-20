package alerts

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/alerts"
)

func GetAllAlerts() ([]models.Alert, error) {
	var err error
	// calling repository
	alerts, err := repository.GetAllAlerts()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving alerts : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving alerts",
		}
	}

	return alerts, nil
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	alert, err := repository.GetAlertById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "alert not found",
			}
		}
		logrus.Errorf("error retrieving alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving usalerter %s", id.String()),
		}
	}

	return alert, err
}

// CreateAgenda crée une nouvelle agenda
func CreateAlert(alert models.Alert) (*models.Alert, error) {
	// Vérifier les champs obligatoires
	if alert.Dest == "" || alert.Msg == "" || alert.TypeA == "" {
		return nil, models.ErrorUnprocessableEntity{
			Message: "Champs obligatoires manquants : Dest, Msg, TypeA",
		}
	}
	// Vérifier que l’agenda existe (appel au repository)
	exists, err := repository.CheckAgendaExists(alert.Agenda)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la vérification de l'agenda : %v", err)
	}
	if !exists {
		return nil, models.ErrorUnprocessableEntity{
			Message: fmt.Sprintf("L'agenda %v n'existe pas", alert.Agenda),
		}
	}
	newAlert, err := repository.CreateAlert(alert)
	if err != nil {
		logrus.Errorf("error creating alert: %s", err.Error())
		return nil, fmt.Errorf("something went wrong while creating alert")
	}
	return newAlert, nil
}

// UpdateAlertService met à jour une alerte existante
func UpdateAlert(id uuid.UUID, alert models.Alert) (*models.Alert, error) {
	// 1️⃣ Vérification des champs obligatoires
	if alert.Dest == "" || alert.Msg == "" || alert.TypeA == "" {
		return nil, models.ErrorUnprocessableEntity{
			Message: "Champs manquants : Dest, Msg, TypeA",
		}
	}
	updatedAlert, err := repository.UpdateAlert(alert)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("alert not found")
		}
		logrus.Errorf("error updating alert %d: %s", id.String(), err.Error())
		return nil, fmt.Errorf("something went wrong while updating alert %d", id)
	}
	return updatedAlert, nil
}
// DeleteAgenda supprime une agenda par son ID
func DeleteAlert(id uuid.UUID) error {
	err := repository.DeleteAlert(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("alert not found")
		}
		logrus.Errorf("error deleting alert %d: %s", id.String(), err.Error())
		return fmt.Errorf("something went wrong while deleting alert %d", id.String())
	}
	return nil
}

// DeleteAllAgendas supprime toutes les agendas
func DeleteAllAlerts() error {
	err := repository.DeleteAllAlerts()
	if err != nil {
		logrus.Errorf("error deleting all alerts: %s", err.Error())
		return fmt.Errorf("something went wrong while deleting all alerts")
	}
	return nil
}


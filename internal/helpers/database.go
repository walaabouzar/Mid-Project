package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)


func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:timetable.db")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	
	// Vérifier que la connexion fonctionne
	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	logrus.Info("✅ Connexion à la base de données réussie")
	return db, nil
	
}
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logrus.Errorf("error closing db : %s", err.Error())
	}
}

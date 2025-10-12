package helpers

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// ErrNoRows est utilisé lorsqu'aucune ligne n'est trouvée dans la BDD
var ErrNoRows = errors.New("no rows found")

// OpenDB ouvre la base SQLite codée en dur
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db") // chemin codé en dur
	if err != nil {
		logrus.Errorf("error opening db: %s", err.Error())
		return nil, err
	}
	db.SetMaxOpenConns(1)
	return db, nil
}

// CloseDB ferme la base
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logrus.Errorf("error closing db : %s", err.Error())
	}
}

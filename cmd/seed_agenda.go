package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	agendas := []struct {
		GroupID string
		ID      string
	}{
		{"M1_G1_langue", "13295"},
		{"M1_G2_langue", "13345"},
		{"M1_G3_langue", "13397"},
		{"M1_G1_option", "7224"},
		{"M1_G2_option", "7225"},
		{"M1_G3_option", "62962"},
		{"M1_G_option", "62090"},
		{"M1_tutorat_L2", "56529"},
	}

	baseURL := "https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=3&calType=ical&nbWeeks=8&displayConfigId=128"

	for _, a := range agendas {
		url := fmt.Sprintf(baseURL, a.ID)
		_, err := db.Exec("INSERT INTO agendas (group_id, ical_url) VALUES (?, ?)", a.GroupID, url)
		if err != nil {
			log.Printf("Erreur insertion %s : %v\n", a.GroupID, err)
		} else {
			fmt.Printf(" Ajouté : %s\n", a.GroupID)
		}
	}

	fmt.Println(" Tous les agendas ont été ajoutés.")
}

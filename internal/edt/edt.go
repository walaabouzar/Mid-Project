package edt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Agenda struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetAgendas(configAPI string) ([]Agenda, error) {
	resp, err := http.Get(configAPI + "/agendas")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("config api returned %d", resp.StatusCode)
	}

	var agendas []Agenda
	if err := json.NewDecoder(resp.Body).Decode(&agendas); err != nil {
		return nil, err
	}

	return agendas, nil
}

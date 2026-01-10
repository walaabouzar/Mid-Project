package scheduler

import (
	"log"

	"github.com/nats-io/nats.go"
)



// InitNatsWithURL initialise la connexion à un NATS central
func InitNatsWithURL(natsURL string) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Erreur connexion NATS:", err)
	}

	js, err = nc.JetStream()
	if err != nil {
		log.Fatal("Erreur JetStream:", err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "COURS",
		Subjects: []string{"COURS.>"},
	})
	if err != nil {
		log.Println("Stream COURS déjà existant ou erreur ignorée")
	}
}

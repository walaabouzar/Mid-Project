package helpers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// FetchCalendar télécharge le contenu iCal d’un agenda donné à partir de son ID
func FetchCalendar(agendaID string) (string, error) {
	url := fmt.Sprintf(
		"https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=3&calType=ical&nbWeeks=8&displayConfigId=128",
		agendaID,
	)

	logrus.Infof(" Téléchargement du calendrier pour AgendaID=%s ...", agendaID)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("erreur lors du téléchargement du calendrier : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("réponse HTTP inattendue : %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lecture du corps de la réponse : %v", err)
	}

	logrus.Infof("✅ Téléchargement réussi pour AgendaID=%s (%d octets)", agendaID, len(body))
	return string(body), nil
}
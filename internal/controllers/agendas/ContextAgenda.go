package agendas

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ContextAgenda récupère l'ID d'un agenda depuis l'URL et le met dans le contexte
func ContextAgenda(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Impossible de parser l'id (%s) en int", idStr), http.StatusUnprocessableEntity)
			return
		}
		ctx := context.WithValue(r.Context(), "agendaID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

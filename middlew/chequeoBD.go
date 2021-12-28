package middlew

import (
	"net/http"

	"github.com/ilopezor/proyecto-twitter-back/bd"
)

/*CheckBD es el middlew que me permite conocer el estado de la BD */
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

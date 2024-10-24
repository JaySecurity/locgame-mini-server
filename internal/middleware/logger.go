package middleware

import (
	"locgame-mini-server/pkg/log"
	"net/http"
)

func Log(handler http.HandlerFunc) http.HandlerFunc {
	return EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("method: %s route: %s ", r.Method, r.URL.Path)
		// Pass control back to the handler
		handler.ServeHTTP(w, r)
	}))
}

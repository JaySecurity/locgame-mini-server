package middleware

import (
	"locgame-mini-server/pkg/log"
	"net/http"
)

// TODO: load origins from config file.
var allowedOrigins = map[string]bool{
	"http://localhost:5173": true,
	"":                      true,
}

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers

		origin := r.Header.Get("Origin")

		if ok := allowedOrigins[origin]; !ok {
			log.Infof("Origin not allowed by CORS: %v", origin)
			w.WriteHeader(http.StatusForbidden) // Forbidden
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin) // Allow all origins (you can restrict this)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		// If it's a preflight OPTIONS request, respond with OK status
		if r.Method == http.MethodOptions {
			log.Debug("MethodOptions", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			return
		}

		// For other requests, pass to the next handler
		next(w, r)
	})
}

package middleware

import (
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/log"
	"net/http"
)

type Middleware struct {
	config         *config.Config
	allowedOrigins map[string]bool
}

func NewMiddleWare(cfg *config.Config) *Middleware {

	m := new(Middleware)
	m.config = cfg
	m.allowedOrigins = make(map[string]bool)

	for _, v := range cfg.AllowedOrigins {
		m.allowedOrigins[v] = true
	}
	m.allowedOrigins["http://192.168.2.42:5173"] = true
	m.allowedOrigins["http://192.168.2.42:5173/"] = true
	log.Debug(m.allowedOrigins)
	return m
}

func (m *Middleware) EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers

		origin := r.Header.Get("Origin")
		log.Debugf("|%s|", origin)

		if ok := m.allowedOrigins[origin]; !ok {
			log.Infof("Origin not allowed by CORS: %v", origin)
			w.WriteHeader(http.StatusForbidden) // Forbidden
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin) // Allow all origins (you can restrict this)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, OPTIONS")
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

func (m *Middleware) Logger(handler http.HandlerFunc) http.HandlerFunc {
	return m.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("method: %s route: %s ", r.Method, r.URL.Path)
		// Pass control back to the handler
		handler.ServeHTTP(w, r)
	}))
}

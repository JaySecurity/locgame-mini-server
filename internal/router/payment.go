package router

import (
	"locgame-mini-server/pkg/log"
	"net/http"
)

func (r *Router) HandlePaymentRoutes() {
	// Get Store Data
	r.Mux.HandleFunc("/payment", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		cookies := req.Cookies()
		for _, cookie := range cookies {
			log.Debug(cookie)
		}
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Payments"))
	})
}

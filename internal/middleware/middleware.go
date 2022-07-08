package middleware

import (
	"log"
	"net/http"
	"time"
)

// EventLogger provided logging middleware
func EventLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Method: %s URI: %s Time: %s", r.Method, r.RequestURI, now)
		},
	)
}

package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			" %-5s %-20s %10s",
			r.Method,
			r.URL.RequestURI(),
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

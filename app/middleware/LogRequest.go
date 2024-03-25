package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	file, err := os.OpenFile("logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.SetOutput(file)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"%s %s (%s)",
			r.Method,
			r.URL.RequestURI(),
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

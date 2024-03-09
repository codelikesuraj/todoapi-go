package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			mux.CurrentRoute(r).GetName(),
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

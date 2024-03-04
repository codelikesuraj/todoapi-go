package helpers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func AcceptsJson(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json")
}

func JsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func RenderTemplate(w http.ResponseWriter, filenames []string, data any) {
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

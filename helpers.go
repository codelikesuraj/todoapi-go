package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

func acceptsJson(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json")
}

func jsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func renderTemplate(w http.ResponseWriter, filenames []string, data any) {
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}

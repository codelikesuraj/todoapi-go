package utils

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AcceptsJson(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json")
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ParsePage(r *http.Request) (int, int) {
	rows := 5
	offset, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if offset < 2 {
		return 0, rows
	}

	return (offset - 1) * rows, rows
}

func RenderTemplate(w http.ResponseWriter, filenames []string, data interface{}) {
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func acceptsJson(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json")
}

func jsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

	if statusCode != 0 {
		w.WriteHeader(statusCode)
	}
}

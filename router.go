package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.MethodNotAllowedHandler = Logger(http.HandlerFunc(CustomMethodNotAllowed), "ErrorRouteNotFound")
	router.NotFoundHandler = Logger(http.HandlerFunc(CustomNotFound), "ErrorMethodNotAllowed")
	for _, route := range routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(Logger(route.HandlerFunc, route.Name))
	}
	return router
}

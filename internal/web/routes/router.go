package routes

import (
	"net/http"
	"todoapi/internal/helpers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.MethodNotAllowedHandler = helpers.Logger(http.HandlerFunc(CustomMethodNotAllowed), "ErrorMethodNotAllowed")
	router.NotFoundHandler = helpers.Logger(http.HandlerFunc(CustomNotFound), "ErrorRouteNotFound")
	
	for _, route := range Routes() {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(helpers.Logger(route.HandlerFunc, route.Name))
	}
	
	return router
}

func CustomNotFound(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "we no get am"}
	if helpers.AcceptsJson(r) {
		helpers.JsonResponse(w, msg)
		return
	}

	helpers.RenderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/errors/404.html",
		},
		msg,
	)
}

func CustomMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "We no like this your manner of approach"}
	if helpers.AcceptsJson(r) {
		helpers.JsonResponse(w, msg)
		return
	}

	helpers.RenderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/errors/405.html",
		},
		msg,
	)
}

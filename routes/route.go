package routes

import (
	"net/http"
	todo_controller "todoapi/app/controllers"
	"todoapi/app/utils"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Name        string
}

func Routes() []Route {
	return []Route{
		/**
		 * +--------------------------+
		 * | ROUTES ARE DEFINED BELOW |
		 * |    |||            |||    |
		 * |    |||            |||    |
		 * |  \ ||| /        \ ||| /  |
		 * |   \|||/          \|||/   |
		 * |    \|/            \|/    |
		 * |     *              *     |
		 * +--------------------------+
		 */
		{"GET", "/todos", todo_controller.Index, "TodoIndex"},
		{"GET", "/todos/create", todo_controller.Create, "TodoCreate"},
		{"POST", "/todos/store", todo_controller.Store, "TodoStore"},
		{"GET", "/todos/completed", todo_controller.ShowCompleted, "TodoStatusCompleted"},
		{"GET", "/todos/pending", todo_controller.ShowPending, "TodoStatusPending"},
		{"GET", "/todos/{id}", todo_controller.ShowById, "TodoShow"},
		{"POST", "/todos/{id}/status/update", todo_controller.ChangeStatus, "TodoStatusChange"},
	}
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.MethodNotAllowedHandler = utils.Logger(http.HandlerFunc(MethodNotAllowedHandler), "ErrorMethodNotAllowed")
	router.NotFoundHandler = utils.Logger(http.HandlerFunc(NotFoundHandler), "ErrorRouteNotFound")

	for _, route := range Routes() {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(utils.Logger(route.HandlerFunc, route.Name))
	}

	return router
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	msg := map[string]string{"message": "We no like this your manner of approach"}
	if utils.AcceptsJson(r) {
		utils.JsonResponse(w, msg)
		return
	}

	utils.RenderTemplate(
		w,
		[]string{
			"./resources/views/layout.html",
			"./resources/views/errors/405.html",
		},
		msg,
	)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	msg := map[string]string{"message": "we no get am"}
	if utils.AcceptsJson(r) {
		utils.JsonResponse(w, msg)
		return
	}

	utils.RenderTemplate(
		w,
		[]string{
			"./resources/views/layout.html",
			"./resources/views/errors/404.html",
		},
		msg,
	)
}

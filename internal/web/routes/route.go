package routes

import (
	"todoapi/internal/controllers"
	"todoapi/internal/models"
)

func Routes() models.Routes {
	return models.Routes{
		/**
		 * routes are defined here
		 */

		models.Route{Name: "TodoIndex", Method: "GET", Pattern: "/todos", HandlerFunc: controllers.Index},
		models.Route{Name: "TodoCreate", Method: "GET", Pattern: "/todos/create", HandlerFunc: controllers.Create},
		models.Route{Name: "TodoStore", Method: "POST", Pattern: "/todos/store", HandlerFunc: controllers.Store},
		models.Route{Name: "TodoStatusCompleted", Method: "GET", Pattern: "/todos/completed", HandlerFunc: controllers.ShowCompleted},
		models.Route{Name: "TodoStatusPending", Method: "GET", Pattern: "/todos/pending", HandlerFunc: controllers.ShowPending},
		models.Route{Name: "TodoShow", Method: "GET", Pattern: "/todos/{id}", HandlerFunc: controllers.ShowById},
	}
}

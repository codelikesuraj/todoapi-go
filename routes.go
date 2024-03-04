package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoCreate", "GET", "/todos/create", TodoCreate},
	Route{"TodoStore", "POST", "/todos/store", TodoStore},
	Route{"TodoStatusCompleted", "GET", "/todos/completed", TodoShowCompleted},
	Route{"TodoStatusPending", "GET", "/todos/pending", TodoShowPending},
	Route{"TodoShow", "GET", "/todos/{id}", TodoShow},
}

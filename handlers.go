package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CustomNotFound(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "we no get am"}
	w.WriteHeader(http.StatusNotFound)

	if acceptsJson(r) {
		jsonResponse(w, 0, msg)
		return
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/errors/404.html")
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(w, msg); err != nil {
		panic(err)
	}
}

func CustomMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "We no like this your manner of approach"}
	w.WriteHeader(http.StatusMethodNotAllowed)

	if acceptsJson(r) {
		jsonResponse(w, 0, msg)
		return
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/errors/405.html")
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(w, msg); err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	if acceptsJson(r) {
		jsonResponse(w, http.StatusOK, map[string]string{
			"message": "Welcome!!! This is a response to a json request",
		})
		return
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/homepage.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = tmpl.Execute(w, nil); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		jsonResponse(w, http.StatusUnprocessableEntity, map[string]string{"message": "unprocessable entity"})
		return
	}

	t, err := RepoCreateTodo(todo)
	if err != nil {
		jsonResponse(w, http.StatusUnprocessableEntity, map[string]string{"message": fmt.Sprint(err)})
		return
	}

	jsonResponse(w, http.StatusCreated, t)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	if acceptsJson(r) {
		jsonResponse(w, http.StatusOK, repoTodos)
		return
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/todos/index.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = tmpl.Execute(w, map[string]Todos{"todos": repoTodos}); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := RepoFindTodo(id)
	if err != nil {
		msg := map[string]string{"message": fmt.Sprint(err)}
		statusCode := http.StatusNotFound

		if acceptsJson(r) {
			jsonResponse(w, statusCode, msg)
		} else {
			tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/errors/404.html")
			if err != nil {
				panic(err)
			}

			w.WriteHeader(statusCode)
			if err = tmpl.Execute(w, msg); err != nil {
				panic(err)
			}
		}
		return
	}

	if acceptsJson(r) {
		jsonResponse(w, http.StatusOK, todo)
		return
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/todos/show.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = tmpl.Execute(w, map[string]Todo{"todo": todo}); err != nil {
		panic(err)
	}
}

func TodoShowCompleted(w http.ResponseWriter, r *http.Request) {
	TodoShowByStatus(w, r, true)
}

func TodoShowPending(w http.ResponseWriter, r *http.Request) {
	TodoShowByStatus(w, r, false)
}

func TodoShowByStatus(w http.ResponseWriter, r *http.Request, status bool) {
	todos := RepoFindTodoByStatus(status)
	if acceptsJson(r) {
		jsonResponse(w, http.StatusOK, todos)
		return
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/todos/index.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = tmpl.Execute(w, map[string]Todos{"todos": todos}); err != nil {
		panic(err)
	}
}

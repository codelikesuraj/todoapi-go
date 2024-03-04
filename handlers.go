package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CustomNotFound(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "we no get am"}
	if acceptsJson(r) {
		jsonResponse(w, msg)
		return
	}

	renderTemplate(
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
	if acceptsJson(r) {
		jsonResponse(w, msg)
		return
	}

	renderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/errors/405.html",
		},
		msg,
	)
}

func Index(w http.ResponseWriter, r *http.Request) {
	if acceptsJson(r) {
		jsonResponse(w, map[string]string{
			"message": "Welcome!!! This is a response to a json request",
		})
		return
	}

	renderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/homepage.html",
		},
		nil,
	)
}

func TodoStore(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	if acceptsJson(r) {
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			jsonResponse(w, map[string]string{"message": "unprocessable entity"})
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}
		todo.Name = r.FormValue("name")
	}

	t, err := RepoCreateTodo(todo)

	if acceptsJson(r) {
		if err != nil {
			jsonResponse(w, map[string]string{"message": fmt.Sprint(err)})
		} else {
			jsonResponse(w, t)
		}

		return
	}

	http.Redirect(w, r, "/todos/create?message=todo created successfully", http.StatusMovedPermanently)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	if acceptsJson(r) {
		CustomNotFound(w, r)
		return
	}

	renderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/todos/create.html",
		},
		map[string]string{
			"message": r.URL.Query().Get("message"),
		},
	)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := RepoGetAll()

	if acceptsJson(r) {
		jsonResponse(w, todos)
		return
	}

	renderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/todos/index.html",
		},
		map[string]any{
			"todos":     todos,
			"pageTitle": "Todos - List",
		},
	)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := RepoFindTodoById(id)
	if err != nil {
		msg := map[string]string{"message": fmt.Sprint(err)}

		if acceptsJson(r) {
			jsonResponse(w, msg)
			return
		} else {
			renderTemplate(
				w,
				[]string{
					"./templates/layout.html",
					"./templates/errors/404.html",
				},
				msg,
			)
		}
		return
	}

	if acceptsJson(r) {
		jsonResponse(w, todo)
		return
	}

	renderTemplate(w,
		[]string{
			"./templates/layout.html",
			"./templates/todos/show.html",
		},
		map[string]any{
			"todo":      todo,
			"pageTitle": fmt.Sprint("Todos - ", todo.Id),
		},
	)
}

func TodoShowCompleted(w http.ResponseWriter, r *http.Request) {
	TodoShowByStatus(w, r, true)
}

func TodoShowPending(w http.ResponseWriter, r *http.Request) {
	TodoShowByStatus(w, r, false)
}

func TodoShowByStatus(w http.ResponseWriter, r *http.Request, status bool) {
	pageTitle := "Todo - Completed"
	if !status {
		pageTitle = "Todo - Pending"
	}
	todos := RepoFindTodoByStatus(status)
	if acceptsJson(r) {
		jsonResponse(w, todos)
		return
	}

	renderTemplate(
		w,
		[]string{
			"./templates/layout.html",
			"./templates/todos/index.html",
		},
		map[string]any{
			"todos":     todos,
			"pageTitle": pageTitle,
		},
	)
}

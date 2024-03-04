package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todoapi/internal/helpers"
	"todoapi/internal/models"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	todos := todo.FetchAll()

	if helpers.AcceptsJson(r) {
		helpers.JsonResponse(w, todos)
		return
	}

	helpers.RenderTemplate(
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

func Store(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if helpers.AcceptsJson(r) {
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			helpers.JsonResponse(w, map[string]string{"message": "unprocessable entity"})
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}
		todo.Name = r.FormValue("name")
	}

	t, err := todo.Create(todo)

	if helpers.AcceptsJson(r) {
		if err != nil {
			helpers.JsonResponse(w, map[string]string{"message": fmt.Sprint(err)})
		} else {
			helpers.JsonResponse(w, t)
		}

		return
	}

	http.Redirect(w, r, "/todos/create?message=todo created successfully", http.StatusMovedPermanently)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if helpers.AcceptsJson(r) {
		http.NotFound(w, r)
		return
	}

	helpers.RenderTemplate(
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

func ShowById(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := t.FindById(id)
	if err != nil {
		msg := map[string]string{"message": fmt.Sprint(err)}

		if helpers.AcceptsJson(r) {
			helpers.JsonResponse(w, msg)
			return
		} else {
			helpers.RenderTemplate(
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

	if helpers.AcceptsJson(r) {
		helpers.JsonResponse(w, todo)
		return
	}

	helpers.RenderTemplate(w,
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

func ShowCompleted(w http.ResponseWriter, r *http.Request) {
	ShowByStatus(w, r, true)
}

func ShowPending(w http.ResponseWriter, r *http.Request) {
	ShowByStatus(w, r, false)
}

func ShowByStatus(w http.ResponseWriter, r *http.Request, status bool) {
	var t models.Todo
	pageTitle := "Todo - Completed"
	if !status {
		pageTitle = "Todo - Pending"
	}
	todos := t.FindByStatus(status)
	if helpers.AcceptsJson(r) {
		helpers.JsonResponse(w, todos)
		return
	}

	helpers.RenderTemplate(
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

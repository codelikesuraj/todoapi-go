package todo_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todoapi/app/models"
	"todoapi/app/utils"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	todos := models.FetchAllTodos()

	if utils.AcceptsJson(r) {
		utils.JsonResponse(w, todos)
		return
	}

	utils.RenderTemplate(
		w,
		[]string{
			"./resources/views/layout.html",
			"./resources/views/todos/index.html",
		},
		map[string]any{
			"todos":     todos,
			"pageTitle": "Todos - List",
		},
	)
}

func Store(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if utils.AcceptsJson(r) {
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			utils.JsonResponse(w, map[string]string{"message": "unprocessable entity"})
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}
		todo.Name = r.FormValue("name")
	}

	t, err := models.CreateTodo(todo)

	if utils.AcceptsJson(r) {
		if err != nil {
			utils.JsonResponse(w, map[string]string{"message": fmt.Sprint(err)})
		} else {
			utils.JsonResponse(w, t)
		}
		return
	}

	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/todos/create?error=%s", err), http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/todos/create?message=todo created successfully", http.StatusMovedPermanently)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if utils.AcceptsJson(r) {
		w.WriteHeader(http.StatusNotFound)
		utils.JsonResponse(w, map[string]string{"message": "unavailable for api routes"})
		return
	}

	utils.RenderTemplate(
		w,
		[]string{
			"./resources/views/layout.html",
			"./resources/views/todos/create.html",
		},
		map[string]string{
			"message": r.URL.Query().Get("message"),
			"error":   r.URL.Query().Get("error"),
		},
	)
}

func ShowById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := models.FindTodoById(id)
	if err != nil {
		msg := map[string]string{"message": fmt.Sprint(err)}

		if utils.AcceptsJson(r) {
			utils.JsonResponse(w, msg)
			return
		} else {
			utils.RenderTemplate(
				w,
				[]string{
					"./resources/views/layout.html",
					"./resources/views/errors/404.html",
				},
				msg,
			)
		}
		return
	}

	if utils.AcceptsJson(r) {
		utils.JsonResponse(w, todo)
		return
	}

	utils.RenderTemplate(w,
		[]string{
			"./resources/views/layout.html",
			"./resources/views/todos/show.html",
		},
		map[string]any{
			"todo":      todo,
			"pageTitle": fmt.Sprint("Todos - ", todo.Id),
			"message":   r.URL.Query().Get("message"),
			"error":     r.URL.Query().Get("error"),
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
	pageTitle := "Todo - Completed"
	if !status {
		pageTitle = "Todo - Pending"
	}
	todos := models.FetchTodosByStatus(status)
	if utils.AcceptsJson(r) {
		utils.JsonResponse(w, todos)
		return
	}

	utils.RenderTemplate(
		w,
		[]string{
			"./resources/views/layout.html",
			"./resources/views/todos/index.html",
		},
		map[string]any{
			"todos":     todos,
			"pageTitle": pageTitle,
		},
	)
}

func ChangeStatus(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := models.FindTodoById(id)
	if err != nil {
		msg := map[string]string{"message": fmt.Sprint(err)}

		if utils.AcceptsJson(r) {
			utils.JsonResponse(w, msg)
			return
		} else {
			utils.RenderTemplate(
				w,
				[]string{
					"./resources/views/layout.html",
					"./resources/views/errors/404.html",
				},
				msg,
			)
		}
		return
	}

	todo, err = models.ChangeTodoStatus(todo)

	if utils.AcceptsJson(r) {
		if err != nil {
			utils.JsonResponse(w, map[string]string{"message": fmt.Sprint(err)})
		} else {
			utils.JsonResponse(w, todo)
		}

		return
	}

	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/todos/%d?error=%s", todo.Id, err), http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/todos/%d?message=todo updated successfully", todo.Id), http.StatusMovedPermanently)
	}
}

package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"todoapi/app/utils"

	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int       `field:"id" json:"id"`
	Name      string    `field:"name" json:"name"`
	Completed bool      `field:"completed" json:"completed"`
	Due       time.Time `field:"due" json:"due"`
}

type Todos []Todo

var db *sql.DB

func init() {
	godotenv.Load()
	db, _ = utils.GetDatabaseConnection()
}

func FetchAllTodos() Todos {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}

	var (
		todo  Todo
		todos Todos
	)

	for rows.Next() {
		if err := rows.Scan(&todo.Id, &todo.Completed, &todo.Due, &todo.Name); err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	if len(todos) > 0 {
		return todos
	}

	return Todos{}
}

func CreateTodo(todo Todo) (Todo, error) {
	if len(todo.Name) < 1 {
		return Todo{}, fmt.Errorf("name is required")
	}

	stmt, err := db.Prepare("INSERT INTO todos (todo, completed, due) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	if todo.Due.Equal(time.Time{}) {
		todo.Due = time.Now().Add(time.Hour)
	}

	result, err := stmt.Exec(todo.Name, todo.Completed, todo.Due)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return FindTodoById(int(id))
}

func FindTodoById(id int) (Todo, error) {
	var todo Todo

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	err := row.Scan(&todo.Id, &todo.Completed, &todo.Due, &todo.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, fmt.Errorf("todo not found")
		}

		log.Fatal(err)
	}
	return todo, nil
}

func FetchTodosByStatus(status bool) Todos {
	rows, err := db.Query("SELECT * FROM todos WHERE completed = ?", status)
	if err != nil {
		log.Fatal(err)
	}

	var (
		todo  Todo
		todos Todos
	)

	for rows.Next() {
		if err := rows.Scan(&todo.Id, &todo.Completed, &todo.Due, &todo.Name); err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	if len(todos) > 0 {
		return todos
	}

	return Todos{}
}

func ChangeTodoStatus(t Todo) (Todo, error) {
	stmt, err := db.Prepare("UPDATE todos SET completed = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(!t.Completed, t.Id)
	if err != nil {
		log.Fatal(err)
	}

	updatedTodo, err := FindTodoById(t.Id)
	if err != nil {
		return t, err
	}

	if updatedTodo.Completed == t.Completed {
		return t, fmt.Errorf("error updating todo")
	}

	return updatedTodo, nil
}

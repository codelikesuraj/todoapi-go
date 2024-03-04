package models

import (
	"database/sql"
	"fmt"
	"time"
	"todoapi/internal/helpers"
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
	db, _ = helpers.GetDatabaseConnection()
}

func (t Todo) FetchAll() Todos {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
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

func (t Todo) Create(todo Todo) (Todo, error) {
	if len(todo.Name) < 1 {
		return Todo{}, fmt.Errorf("name is required")
	}

	stmt, err := db.Prepare("INSERT INTO todos (todo, completed, due) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if todo.Due.Equal(time.Time{}) {
		todo.Due = time.Now().Add(time.Hour)
	}

	result, err := stmt.Exec(todo.Name, todo.Completed, todo.Due)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return t.FindById(int(id))
}

func (t Todo) FindById(id int) (Todo, error) {
	var todo Todo

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	err := row.Scan(&todo.Id, &todo.Completed, &todo.Due, &todo.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, fmt.Errorf("todo not found")
		}

		panic(err)
	}
	return todo, nil
}

func (t Todo) FindByStatus(status bool) Todos {
	rows, err := db.Query("SELECT * FROM todos WHERE completed = ?", status)
	if err != nil {
		panic(err)
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

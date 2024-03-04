package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func init() {

	fmt.Println("Setting up database...")

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/todoapigo?parseTime=true")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 500)

	fmt.Println("Database setup complete")
}

func RepoGetAll() Todos {
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

func RepoCreateTodo(t Todo) (Todo, error) {
	if len(t.Name) < 1 {
		return Todo{}, fmt.Errorf("name is required")
	}

	stmt, err := db.Prepare("INSERT INTO todos (todo, completed, due) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if t.Due.Equal(time.Time{}) {
		t.Due = time.Now().Add(time.Hour)
	}

	result, err := stmt.Exec(t.Name, t.Completed, t.Due)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return RepoFindTodoById(int(id))
}

func RepoDestroyTodo(id int) error {
	todo, err := RepoFindTodoById(id)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(todo.Id)
	if err != nil {
		return err
	}

	return nil
}

func RepoFindTodoById(id int) (Todo, error) {
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

func RepoFindTodoByStatus(status bool) Todos {
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

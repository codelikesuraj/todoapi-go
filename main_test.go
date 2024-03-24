package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todoapi/app/models"
	"todoapi/app/utils"
	"todoapi/routes"
)

func TestListTodos(t *testing.T) {
	numOfRows := 7

	refreshDatabase()
	seedDatabase(numOfRows)

	req, err := http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()

	routes.Router().ServeHTTP(rr, req)

	t.Run("check status OK - 200", func(t *testing.T) {
		want := http.StatusOK
		got := rr.Code
		if want != got {
			t.Errorf("incorrect status code - expected '%v', got '%v'", want, got)
		}
	})

	t.Run("check number of todos", func(t *testing.T) {
		got, want := 0, numOfRows
		rows, err := dbConn().Query("SELECT * FROM todos LIMIT ?", want)
		if err != nil {
			t.Errorf(err.Error())
		}

		for rows.Next() {
			got++
		}

		if want != got {
			t.Errorf("incorrect rows returned - expected '%v' rows, got '%v'", want, got)
		}
	})
}

func TestStoreTodo(t *testing.T) {
	refreshDatabase()

	todo := models.Todo{Name: "Todo X"}
	body, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf(err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, "/todos/store", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()

	routes.Router().ServeHTTP(rr, req)

	t.Run("check status CREATED - 201", func(t *testing.T) {
		want := http.StatusCreated
		got := rr.Code
		if want != got {
			t.Errorf("incorrect status code - expected '%v', got '%v'", want, got)
		}
	})

	t.Run("check response body", func(t *testing.T) {
		var got models.Todo
		want := todo
		if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
			t.Fatalf(err.Error())
		}

		if want.Name != got.Name {
			t.Errorf("todo name not found in response body - expected %q got %q", want.Name, got.Name)
		}
	})
}

func TestFetchTodo(t *testing.T) {
	refreshDatabase()
	seedDatabase(1)

	t.Run("check status OK - 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/todos/1", nil)
		if err != nil {
			t.Fatal(err.Error())
		}
		req.Header.Set("Accept", "application/json")

		rr := httptest.NewRecorder()

		routes.Router().ServeHTTP(rr, req)

		want := http.StatusOK
		got := rr.Code
		if want != got {
			t.Errorf("incorrect status code - expected '%v', got '%v'", want, got)
		}
	})

	t.Run("check status NOT FOUND - 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/todos/999", nil)
		if err != nil {
			t.Fatal(err.Error())
		}
		req.Header.Set("Accept", "application/json")

		rr := httptest.NewRecorder()

		routes.Router().ServeHTTP(rr, req)

		want := http.StatusNotFound
		got := rr.Code
		if want != got {
			t.Errorf("incorrect status code - expected '%v', got '%v'", want, got)
		}
	})
}

func refreshDatabase() {
	_, err := dbConn().Exec("DELETE FROM todos")
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = dbConn().Exec("ALTER TABLE todos AUTO_INCREMENT = 1")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func seedDatabase(rows int) {
	var (
		dues  []interface{}
		query string = "INSERT INTO todos (due, todo) VALUES "
	)

	for i := 0; i < rows; i++ {
		dues = append(dues, time.Now().Add(time.Hour))
		query += fmt.Sprintf("(?, 'Todo %v')", i)
		if i < rows-1 {
			query += ","
		}
	}

	stmt, err := dbConn().Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = stmt.Exec(dues...)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func dbConn() *sql.DB {
	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}

package main

import (
	"fmt"
	"time"
)

var (
	repoId    int
	repoTodos Todos

	ErrTodoNotFound = "todo not found"
)

func init() {
	fmt.Println("Setting up database...")
	time.Sleep(time.Second)
	fmt.Println("Database setup complete")

	RepoCreateTodo(Todo{Name: "Host meetup", Due: time.Now().Add(time.Hour)})
	RepoCreateTodo(Todo{Name: "Publish article", Due: time.Now().Add(time.Hour)})
	RepoCreateTodo(Todo{Name: "Write presentation", Completed: true, Due: time.Now()})
}

func RepoCreateTodo(t Todo) (Todo, error) {
	if len(t.Name) < 1 {
		return Todo{}, fmt.Errorf("name is required")
	}
	repoId += 1
	t.Id = repoId
	repoTodos = append(repoTodos, t)
	return t, nil
}

func RepoDestroyTodo(id int) error {
	for i, t := range repoTodos {
		if t.Id == id {
			repoTodos = append(repoTodos[:i], repoTodos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%s with id: %d", ErrTodoNotFound, id)
}

func RepoFindTodo(id int) (Todo, error) {
	for _, todo := range repoTodos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return Todo{}, fmt.Errorf("%s with id: %d", ErrTodoNotFound, id)
}

func RepoFindTodoByStatus(status bool) Todos {
	var todos Todos

	for _, todo := range repoTodos {
		if todo.Completed == status {
			todos = append(todos, todo)
		}
	}

	return todos
}

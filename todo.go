package main

import "time"

type Todo struct {
	Id        int       `field:"id" json:"id"`
	Name      string    `field:"name" json:"name"`
	Completed bool      `field:"completed" json:"completed"`
	Due       time.Time `field:"due" json:"due"`
}

type Todos []Todo

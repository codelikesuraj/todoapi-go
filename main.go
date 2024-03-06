package main

import (
	"fmt"
	"log"
	"net/http"
	"todoapi/routes"
)

const SERVER_PORT = "8888"

func main() {
	fmt.Printf("Server running on port: %s\n\n", SERVER_PORT)
	log.Fatal(http.ListenAndServe(":"+SERVER_PORT, routes.Router()))
}

package main

import (
	"fmt"
	"net/http"

	"todoapi/internal/web/routes"
)

const SERVER_PORT = "8888"

func main() {
	fmt.Printf("Server running on port: %s\n\n", SERVER_PORT)

	panic(http.ListenAndServe(":"+SERVER_PORT, routes.Router()))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todoapi/routes"
)

var port = os.Getenv("SERVER_PORT")

func main() {
	fmt.Printf("Server running on port: %s\n\n", port)
	log.Fatal(http.ListenAndServe(":"+port, routes.Router()))
}

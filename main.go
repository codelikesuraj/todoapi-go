package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8888"

func main() {
	fmt.Printf("Server running on port: %s\n\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, NewRouter()))
}

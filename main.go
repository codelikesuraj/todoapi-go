package main

import (
	"fmt"
	"net/http"
)

const PORT = "8888"

func main() {
	fmt.Printf("Server running on port: %s\n\n", PORT)
	panic(http.ListenAndServe(":"+PORT, NewRouter()))
}

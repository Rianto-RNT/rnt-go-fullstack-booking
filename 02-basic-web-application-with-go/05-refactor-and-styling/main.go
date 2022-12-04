package main

import (
	"fmt"
	"net/http"
)

// NOTE : run the main.go using = $ go run .
// Port
const PORT = ":8080"

// @desc    Main Application
// @access  Private
func main() {
	// Routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Port listening
	fmt.Println(fmt.Sprintf("Application is running in port: %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}

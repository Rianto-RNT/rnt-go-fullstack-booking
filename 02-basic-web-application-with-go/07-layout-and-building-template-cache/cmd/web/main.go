package main

import (
	"fmt"
	"net/http"
	"rnt-go-booking/02-basic-web-application-with-go/06-go-modules-and-use-packages/pkg/handlers"
)

// NOTE : run the main.go using = $ go run .
// Port
const PORT = ":8080"

// @desc    Main Application
// @access  Private
func main() {
	// Routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// Port listening
	fmt.Println(fmt.Sprintf("Application is running in port: %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}

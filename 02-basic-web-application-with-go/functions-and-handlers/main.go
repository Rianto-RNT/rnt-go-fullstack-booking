package main

import (
	"fmt"
	"net/http"
)

// Port
const PORT = ":8080"

// @desc    Show Home Page
// @route   GET /
// @access  Public
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// @desc    Show About Page
// @route   GET /about
// @access  Public
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("About Page | addValues result: %d", sum))
}

// @desc    add two integers and return the sum
// @access  Private
func addValues(x, y int) int {
	return x + y
}

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
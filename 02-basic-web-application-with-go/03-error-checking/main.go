package main

import (
	"errors"
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

func Devide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f\n", 100.0, 0.0, f)) // change value here
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

// @desc    Main Application
// @access  Private
func main() {
	// Routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Devide)

	// Port listening
	fmt.Println(fmt.Sprintf("Application is running in port: %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}

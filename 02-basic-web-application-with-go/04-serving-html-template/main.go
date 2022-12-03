package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Port
const PORT = ":8080"

// @desc    Show Home Page
// @route   GET /
// @access  Public
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// @desc    Show About Page
// @route   GET /about
// @access  Public
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
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

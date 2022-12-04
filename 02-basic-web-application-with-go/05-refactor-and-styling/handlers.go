package main

import (
	"net/http"
)

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

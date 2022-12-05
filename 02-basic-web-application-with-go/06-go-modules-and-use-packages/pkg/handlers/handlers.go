package handlers

import (
	"net/http"
	"rnt-go-booking/02-basic-web-application-with-go/06-go-modules-and-use-packages/pkg/render"
)

// @desc    Show Home Page
// @route   GET /
// @access  Public
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// @desc    Show About Page
// @route   GET /about
// @access  Public
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

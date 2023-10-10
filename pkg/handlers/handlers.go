package handlers

import (
	"net/http"

	"github.com/maher-naija-pro/Go-Generic-Web-Applications/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
	return
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
	return
}

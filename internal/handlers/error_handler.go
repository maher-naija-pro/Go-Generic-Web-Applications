package handlers

import (
	"net/http"
	"web_server/internal/models"
	"web_server/internal/render"
)

// login is the handler for the login page
func (m *Repository) Error_Show(w http.ResponseWriter, r *http.Request) {
	
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}





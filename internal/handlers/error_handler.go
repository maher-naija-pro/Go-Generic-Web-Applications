package handlers

import (
	"net/http"
	"web_server/internal/models"
	"web_server/internal/render"
)

// login is the handler for the login page
func (m *Repository) Error_Show(w http.ResponseWriter, r *http.Request) {
	
	stringMap := make(map[string]string)
	stringMap["error"] = m.App.Session.GetString(r.Context(), "error")
	render.RenderTemplate(w, r ,"error.page.tmpl", &models.TemplateData{StringMap: stringMap})
}





package handlers

import (
	"net/http"
	"web_server/internal/models"
	"web_server/internal/render"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r , "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

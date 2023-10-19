package handlers

import (
	"log"
	"net/http"
)

// Login is the handler for the maher page
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	//Dump_req (r)
	m.App.Session.Destroy(r.Context())
	log.Println("logout")

	http.Redirect(w, r, "", http.StatusSeeOther)
}

package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"web_server/internal/models"
	"web_server/internal/render"
)

// login is the handler for the login page
func (m *Repository) Login_Show(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}




// Login is the handler for the maher page
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	r.ParseForm()
	
	email :=  r.Form.Get("email")   
	password := r.Form.Get("password")   
     
	

	ID, err := m.DB.AuthUser(email,password)
	if err != nil {
		log.Println(err)
		m.App.Session.Put(r.Context(), "error", "can't insert reservation into database!")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	_ = ID
	
	resp := jsonResponse{
		OK:      true,
		Message: "hello" ,
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

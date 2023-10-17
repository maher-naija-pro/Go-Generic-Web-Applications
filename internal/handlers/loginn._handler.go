package handlers

import (
	"net/http"
	"log"
	"web_server/internal/models"
	"web_server/internal/forms"
	"web_server/internal/render"
)

// login is the handler for the login page
func (m *Repository) Login_Show(w http.ResponseWriter, r *http.Request) {
	
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}




// Login is the handler for the maher page
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if  err != nil {
       log.Println(err )
	}

	email :=  r.Form.Get("email")   
	password := r.Form.Get("password")

	myForm := forms.New(r.PostForm)
	
	myForm.Required("email", password)
	if !myForm.Valid() { //TODO 
	}

	ID, err := m.DB.AuthUser(email,password)
	if err != nil {
		log.Println("cant login") 
		m.App.Session.Put(r.Context(), "error", "can't login")
		http.Redirect(w, r, "login", http.StatusSeeOther)
	}


	m.App.Session.Put(r.Context(),"user_id",ID) 
	m.App.Session.Put(r.Context(),"flash","logged sucessful")
    http.Redirect(w, r, "", http.StatusSeeOther)	

}

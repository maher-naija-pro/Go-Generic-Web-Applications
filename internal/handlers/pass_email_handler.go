package handlers

import (
	"net/http"
	"log"
	"web_server/internal/models"
	"web_server/internal/forms"
	"web_server/internal/render"
)

// login is the handler for the login page
func (m *Repository) Pass_Email_Show(w http.ResponseWriter, r *http.Request) {
	
	render.RenderTemplate(w,r, "pass_email.page.tmpl", &models.TemplateData{})
}




// Login is the handler for the maher page
func (m *Repository) Pass_Email(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if  err != nil {
       log.Println(err )
	}

	email :=  r.Form.Get("email")   

	myForm := forms.New(r.PostForm)
   _ = email	
	myForm.Required("email")

	if !myForm.Has("email") { 
		m.App.Session.Put(r.Context(), "error", "please give email")
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
		}else if  !myForm.IsEmail("email") {
		m.App.Session.Put(r.Context(), "error", "not a valid email address")
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	

//TODO check if mail exeist else ask to subscribe

	//m.App.Session.Put(r.Context(),"user_id",ID) 
    http.Redirect(w, r, "", http.StatusSeeOther)	

}

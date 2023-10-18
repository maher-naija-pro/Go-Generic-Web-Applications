package handlers

import (
	"net/http"
	"log"
	"web_server/internal/models"
	"web_server/internal/forms"
	"web_server/internal/render"
)

// reset pass tempalte
func (m *Repository) Pass_Reset_Show(w http.ResponseWriter, r *http.Request) {
	
	render.RenderTemplate(w,r,
		 "reset_pass.page.tmpl", &models.TemplateData{})
}




// Login is the handler for the maher page
func (m *Repository) Pass_Reset(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if  err != nil {
       log.Println(err )
	}

	password :=  r.Form.Get("email")   

	myForm := forms.New(r.PostForm)
   _ = password 
   
	myForm.Required("password","password2")

    http.Redirect(w, r, "", http.StatusSeeOther)	

}

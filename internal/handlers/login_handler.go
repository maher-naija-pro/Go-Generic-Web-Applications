package handlers

import (
	"fmt"
	"log"
	"net/http"
	"web_server/internal/forms"
	"web_server/internal/models"
	"web_server/internal/render"
)

// login is the handler for the login page
func (m *Repository) Login_Show(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["test"] = "eee"
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Login is the handler for the maher page
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	//Dump_req (r)
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	myForm := forms.New(r.PostForm)

	myForm.Required("email", "password")
	myForm.IsEmail("email")

	if !myForm.Valid() {
		log.Println("not valid")
		render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
			Forms: myForm,
			Error: "",
		})
		return
	}

	ID, err := m.DB.AuthUser(email, password)
	if err != nil {
		log.Println(err)
		m.App.Session.Put(r.Context(), "error", fmt.Sprint(err))
		render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
			Forms: myForm,
			Error: fmt.Sprint(err),
		})
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	log.Println("your id find", ID)

	m.App.Session.Put(r.Context(), "user_id", ID)
	m.App.Session.Put(r.Context(), "flash", "logged sucessful")
	http.Redirect(w, r, "", http.StatusSeeOther)
}

package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"web_server/internal/models"
	"web_server/internal/render"
)

// subscribe is the handler for the subscribe page
func (m *Repository) Subscribe_Show(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "subscribe.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Subscribe(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	r.ParseForm()


	user := models.User{
		
		FirstName:   r.Form.Get("firstname")   ,
		LastName:    r.Form.Get("lastname")   ,
		Email:       r.Form.Get("email")   ,
		Password:   r.Form.Get("password")   ,
 	}


	ID, err := m.DB.InsertUser(user)
	if err != nil {
		log.Println(err)
		m.App.Session.Put(r.Context(), "error", "can't insert reservation into database!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
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
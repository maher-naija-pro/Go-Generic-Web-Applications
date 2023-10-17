package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"net/http/httputil"
	"web_server/internal/config"
	"web_server/internal/models"
	"web_server/internal/render"
	"web_server/internal/driver"
	"web_server/internal/repository"
	"web_server/internal/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}



type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// login is the handler for the login page
func (m *Repository) Login_Show(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}


// subscribe is the handler for the subscribe page
func (m *Repository) Subscribe_Show(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "subscribe.page.tmpl", &models.TemplateData{})
}


func Dump_req (r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
    if err != nil {
        fmt.Printf("REQUESTn")
    }
    fmt.Printf("REQUEST:\n%s", string(reqDump))
return

}

// Login is the handler for the maher page
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	r.ParseForm()

	user := models.User{
		Username:   r.Form.Get("username")   ,
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
	//w.Write([]byte(fmt.Sprintf("the user %s, pass %s", string(user),string(password))))
	
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

func (m *Repository) Subscribe(w http.ResponseWriter, r *http.Request) {
    //Dump_req (r)
	r.ParseForm()


	user := models.User{
		Username:   r.Form.Get("username")   ,
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

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

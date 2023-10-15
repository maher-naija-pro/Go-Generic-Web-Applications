package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"net/http/httputil"
	"web_server/pkg/config"
	"web_server/pkg/models"
	"web_server/pkg/render"
	"web_server/pkg/driver"
	"web_server/pkg/repository"
	"web_server/pkg/repository/dbrepo"
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

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}


type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
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

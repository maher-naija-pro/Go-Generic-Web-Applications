package helpers

import (
	"fmt"
	"web_server/internal/config"
	"net/http/httputil"
	"net/http"	
	
)


func Dump_req (r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
    if err != nil {
        fmt.Printf("REQUESTn")
    }
    fmt.Printf("REQUEST:\n%s", string(reqDump))
return

}

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}


func IsAuth (r *http.Request)  bool {

	exist:= app.Session.Exists(r.Context() , "user_id")

	return exist
}

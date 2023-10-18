package main

import (
	"github.com/justinas/nosurf"
	"net/http"
	"web_server/internal/helpers"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}


func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) 	{

		if !helpers.IsAuth(r) {
            session.Put (r.Context(), "error", "loginfisrt")
            http.Redirect(w, r, "/login", http.StatusSeeOther)	
		    return
          }    
		  next.ServeHTTP(w,r)
	}	)
	}

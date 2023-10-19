package main

import (
	"net/http"
	"web_server/internal/config"
	"web_server/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Post("/login", handlers.Repo.Login)
	// TODO add logout api
	//mux.Post("/logout", handlers.Repo.Login)
	
    mux.Get("/login", handlers.Repo.Login_Show)
	mux.Post("/subscribe",handlers.Repo.Subscribe)
	mux.Get("/error", handlers.Repo.Error_Show)
	mux.Get("/askpass", handlers.Repo.Pass_Email_Show)
	mux.Post("/askpass",handlers.Repo.Pass_Email)
	//mux.Get("/resetpass", handlers.Repo.Pass_Reset_Show)
	mux.Post("/resetpass",handlers.Repo.Pass_Reset)

	mux.Route("/", func (mux chi.Router) {
      mux.Use(Auth)
      mux.Get("/resetpass", handlers.Repo.Pass_Reset_Show)
	})
	// serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))


	return mux
}

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
    mux.Get("/login", handlers.Repo.Login_Show)
	mux.Post("/subscribe",handlers.Repo.Subscribe)
	mux.Get("/subscribe", handlers.Repo.Subscribe_Show)
	// serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

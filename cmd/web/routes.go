package main

import (
	"net/http"

	"github.com/TheKinng96/Go-booking-app/pkg/config"
	"github.com/TheKinng96/Go-booking-app/pkg/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", controllers.Repo.Home)
	mux.Get("/about", controllers.Repo.About)

	return mux
}
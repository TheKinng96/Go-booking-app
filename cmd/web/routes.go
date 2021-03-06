package main

import (
	"net/http"

	"github.com/TheKinng96/Go-booking-app/internal/config"
	"github.com/TheKinng96/Go-booking-app/internal/controllers"
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
	mux.Get("/generals-quarters", controllers.Repo.Generals)
	mux.Get("/majors-suite", controllers.Repo.Majors)

	mux.Get("/search-availability", controllers.Repo.Availability)
	mux.Post("/search-availability", controllers.Repo.PostAvailability)
	mux.Post("/search-availability-json", controllers.Repo.AvailabilityJson)

	mux.Get("/contact", controllers.Repo.Contact)

	mux.Get("/make-reservation", controllers.Repo.Reservation)
	mux.Post("/make-reservation", controllers.Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

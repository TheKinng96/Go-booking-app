package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TheKinng96/Go-booking-app/internal/config"
	"github.com/TheKinng96/Go-booking-app/internal/controllers"
	"github.com/TheKinng96/Go-booking-app/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// Update the env
	app.InProduction = false

	session = scs.New()
	// How long to keep user logged in
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := controllers.NewRepo(&app)
	controllers.NewControllers(repo)
	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

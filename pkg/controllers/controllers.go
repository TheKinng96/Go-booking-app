package controllers

import (
	"fmt"
	"net/http"

	"github.com/TheKinng96/Go-booking-app/models"
	"github.com/TheKinng96/Go-booking-app/pkg/config"
	"github.com/TheKinng96/Go-booking-app/pkg/render"
)

// Repository used by controllers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Create a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the controllers
func NewControllers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	fmt.Println(remoteIp)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Welcome to about page!"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

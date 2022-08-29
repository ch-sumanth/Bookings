package handlers

import (
	"github.com/ch-sumanth/Bookings/pkg/config"
	"github.com/ch-sumanth/Bookings/pkg/models"
	"github.com/ch-sumanth/Bookings/pkg/render"
	"net/http"
)

// TemplateData holds data used by handlers

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello from about route"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

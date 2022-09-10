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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

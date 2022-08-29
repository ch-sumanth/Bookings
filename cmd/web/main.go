package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ch-sumanth/Bookings/pkg/config"
	"github.com/ch-sumanth/Bookings/pkg/handlers"
	"github.com/ch-sumanth/Bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false // set this as true in production server
	session = scs.New()
	session.Lifetime = 5 * time.Minute
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("Cannot create cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port 8080"))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server")
	}
}

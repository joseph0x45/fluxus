package main

import (
	"embed"
	"flag"
	"fluxus/db"
	"fluxus/handler"
	"fluxus/models"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:generate tailwindcss -i static/input.css -o static/styles.css -m

//go:embed templates/*html
var templatesFS embed.FS

//go:embed static/styles.css
var stylesCSS template.CSS

//go:embed static/alpine.js

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))
}

func main() {
	pageData := models.PageData{
		Title:     "",
		StylesCSS: stylesCSS,
	}
	port := flag.String("port", "8080", "The port to start fluxus on")
	flag.Parse()
	r := chi.NewRouter()

	conn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	handler := handler.NewHandler(conn, templates, pageData)
	handler.RegisterRoutes(r)
	server := http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}
	log.Printf("fluxus launched on http://localhost:%s\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

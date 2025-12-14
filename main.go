package main

import (
	"embed"
	"flag"
	"fluxus/db"
	"fluxus/handler"
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
var alpineJS template.JS

//go:embed static/app.js
var appJS template.JS

func init() {
	initTemplates(templatesFS)
}

func main() {
	port := flag.String("port", "8080", "The port to start fluxus on")
	flag.Parse()
	r := chi.NewRouter()

	conn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	handler := handler.NewHandler(conn)
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

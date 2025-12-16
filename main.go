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

//go:embed static/*
var staticFS embed.FS

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		panic(err)
	}
}

func main() {
	port := flag.String("port", "8080", "The port to start fluxus on")
	flag.Parse()
	r := chi.NewRouter()

	conn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("GET /static/*", http.FileServer(http.FS(staticFS)).ServeHTTP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home", nil)
	})

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

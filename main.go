package main

import (
	"context"
	"flag"
	"fluxus/db"
	"fluxus/handler"
	"fluxus/ui"
	"html/template"
	"log"
	"net/http"
)

//go:generate tailwindcss -i assets/vendor/input.css -o assets/vendor/styles.css -m

var templates *template.Template

func init() {

}
func main() {
	port := flag.String("port", "8080", "The port to start fluxus on")
	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ui.Index().Render(ctx, w)
	})

	conn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	handler := handler.NewHandler(conn)
	handler.RegisterRoutes(mux)
	server := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}
	log.Printf("fluxus launched on http://localhost:%s\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

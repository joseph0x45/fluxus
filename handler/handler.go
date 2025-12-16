package handler

import (
	"fluxus/db"
	"html/template"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	conn      *db.Conn
	templates *template.Template
}

func NewHandler(
	conn *db.Conn,
	templates *template.Template,
) *Handler {
	return &Handler{
		conn:      conn,
		templates: templates,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
}

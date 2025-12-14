package handler

import (
	"fluxus/db"
	"fluxus/models"
	"html/template"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	conn      *db.Conn
	templates *template.Template
	pageData  models.PageData
}

func NewHandler(
	conn *db.Conn, templates *template.Template,
	pageData models.PageData,
) *Handler {
	return &Handler{
		conn:      conn,
		templates: templates,
		pageData: pageData,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
}

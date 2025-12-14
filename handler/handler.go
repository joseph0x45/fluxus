package handler

import (
	"fluxus/db"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	conn *db.Conn
}

func NewHandler(
	conn *db.Conn) *Handler {
	return &Handler{
		conn: conn,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
}

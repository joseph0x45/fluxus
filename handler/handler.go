package handler

import (
	"fluxus/db"
	"net/http"
)

type Handler struct {
	conn *db.Conn
}

func NewHandler(conn *db.Conn) *Handler {
	return &Handler{
		conn: conn,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /auth/", h.RenderAuthPage)
}

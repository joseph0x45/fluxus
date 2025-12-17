package handler

import (
	"fluxus/db"
	"fluxus/models"
	"html/template"
	"net/http"
)

type Handler struct {
	conn      *db.Conn
	templates *template.Template
	mode      string
}

func NewHandler(
	conn *db.Conn,
	templates *template.Template,
	mode string,
) *Handler {
	return &Handler{
		conn:      conn,
		templates: templates,
		mode:      mode,
	}
}

func (h *Handler) InReleaseMode() bool {
	return h.mode == "debug"
}

func (h *Handler) getCurrentUser(w http.ResponseWriter, r *http.Request) *models.User {
	currentUser, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return nil
	}
	return currentUser
}

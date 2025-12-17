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

func (h *Handler) RedirectToAccountsPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/accounts", http.StatusSeeOther)
}

func (h *Handler) RenderAccountsPage(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	pageData := models.PageData{
		Title:      "Fluxus | Accounts",
		ActivePage: "accounts",
		Payload: map[string]any{
			"User": currentUser,
		},
	}
	if err := h.templates.ExecuteTemplate(w, "accounts", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

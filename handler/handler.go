package handler

import (
	"fluxus/db"
	"fluxus/models"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
	r.Post("/api/authenticate", h.HandleAuthentication)
	r.With(h.AuthMiddleware).Group(func(r chi.Router) {
		r.Get("/", h.RenderHomePage)
		r.Post("/toggle-safe-mode", h.ToggleSafeMode)
	})
}

func (h *Handler) getCurrentUser(w http.ResponseWriter, r *http.Request) *models.User {
	currentUser, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return nil
	}
	return currentUser
}

func (h *Handler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	pageData := models.PageData{
		Title: "Fluxus | Home",
		Payload: map[string]any{
			"User": currentUser,
		},
	}
	if err := h.templates.ExecuteTemplate(w, "home", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

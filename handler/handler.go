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
	env       string
}

func NewHandler(
	conn *db.Conn,
	templates *template.Template,
	env string,
) *Handler {
	return &Handler{
		conn:      conn,
		templates: templates,
		env:       env,
	}
}

func (h *Handler) InDevMode() bool {
	return h.env == ""
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
	r.Post("/api/authenticate", h.HandleAuthentication)
	r.With(h.AuthMiddleware).Group(func(r chi.Router) {
		r.Get("/", h.RenderHomePage)
	})
}

func (h *Handler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	currentUser, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
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

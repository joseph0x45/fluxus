package handler

import (
	"fluxus/models"
	"net/http"
)

func (h *Handler) RenderAuthPage(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "auth", models.PageData{Title: "Fluxus | Authenticate"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

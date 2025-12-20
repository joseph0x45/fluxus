package handler

import (
	"fluxus/models"
	"net/http"
)

func (h *Handler) RenderBucketsPage(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	pageData := models.PageData{
		ActivePage: "buckets",
		Title:      "Fluxus | Buckets",
		Payload: map[string]any{
			"User":    currentUser,
			"Buckets": []models.Tag{},
		},
	}
	if err := h.templates.ExecuteTemplate(w, "buckets", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

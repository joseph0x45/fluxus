package handler

import (
	"fluxus/models"
	"net/http"
)

func (h *Handler) RenderTagsPage(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	pageData := models.PageData{
		ActivePage: "tags",
		Title:      "Fluxus | Tags",
		Payload: map[string]any{
			"User": currentUser,
			"Tags": []models.Tag{},
		},
	}
	if err := h.templates.ExecuteTemplate(w, "tags", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

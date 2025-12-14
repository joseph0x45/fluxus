package handler

import (
	"net/http"
)

func (h *Handler) RenderAuthPage(w http.ResponseWriter, r *http.Request) {
	h.pageData.Title = "Fluxus | Login"
	err := h.templates.ExecuteTemplate(w, "auth_page", h.pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

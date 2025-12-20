package handler

import (
	"fluxus/logger"
	"fluxus/models"
	"net/http"
)

func (h *Handler) ToggleSafeMode(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	err := h.conn.ToggleUserSafeMode(currentUser)
	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusNoContent)
		return
	}
	h.templates.ExecuteTemplate(w, "safe_mode_toggle", currentUser)
}

func (h *Handler) RenderSettingsPage(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	pageData := models.PageData{
		Title:      "Fluxus | Settings",
		ActivePage: "settings",
		Payload: map[string]any{
			"User": currentUser,
		},
	}
	if err := h.templates.ExecuteTemplate(w, "settings", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

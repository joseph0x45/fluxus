package handler

import (
	"fluxus/logger"
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

package handler

import (
	"fluxus/templates"
	"net/http"
)

func (h *Handler) RenderAuthPage(w http.ResponseWriter, r *http.Request) {
	pageData := templates.PageData()
	pageData.PageTitle = "Fluxus | Authenticate"
	if err := templates.AuthPage.Execute(w, pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

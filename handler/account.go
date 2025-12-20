package handler

import (
	"fluxus/models"
	"net/http"
)

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

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {}

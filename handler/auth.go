package handler

import (
	"context"
	"fluxus/ui/pages"
	"net/http"
)

func (h *Handler) RenderAuthPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.AuthPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

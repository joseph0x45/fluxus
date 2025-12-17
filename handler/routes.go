package handler

import "github.com/go-chi/chi/v5"

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
	r.Post("/api/authenticate", h.HandleAuthentication)
	r.With(h.AuthMiddleware).Group(func(r chi.Router) {
		r.Post("/logout", h.Logout)
		r.Get("/", h.RedirectToAccountsPage)
		r.Get("/accounts", h.RenderAccountsPage)
		r.Post("/toggle-safe-mode", h.ToggleSafeMode)
		r.Get("/tags", h.RenderTagsPage)
	})
}

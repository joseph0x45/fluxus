package handler

import "github.com/go-chi/chi/v5"

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/auth", h.RenderAuthPage)
	r.Post("/api/authenticate", h.HandleAuthentication)
	r.With(h.AuthMiddleware).Group(func(r chi.Router) {
		r.Post("/logout", h.Logout)
		r.Post("/toggle-safe-mode", h.ToggleSafeMode)

		r.Get("/", h.RedirectToAccountsPage)
		r.Get("/accounts", h.RenderAccountsPage)
		r.Get("/tags", h.RenderTagsPage)
		r.Get("/settings", h.RenderSettingsPage)
		r.Get("/buckets", h.RenderBucketsPage)

		r.Post("/tags", h.CreateTag)
		r.Delete("/tags/{id}", h.DeleteTag)
	})
}

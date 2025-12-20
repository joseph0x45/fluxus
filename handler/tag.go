package handler

import (
	"fluxus/logger"
	"fluxus/models"
	"net/http"

	"github.com/oklog/ulid/v2"
)

func (h *Handler) RenderTagsPage(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	tags, err := h.conn.GetUserTags(currentUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageData := models.PageData{
		ActivePage: "tags",
		Title:      "Fluxus | Tags",
		Payload: map[string]any{
			"User": currentUser,
			"Tags": tags,
		},
	}
	if err := h.templates.ExecuteTemplate(w, "tags", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {
	currentUser := h.getCurrentUser(w, r)
	if currentUser == nil {
		return
	}
	if err := r.ParseForm(); err != nil {
		logger.Err("Failed to parse form", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong! Try again"))
		return
	}
	tagName := r.FormValue("name")
	inserted, err := h.conn.InsertTag(&models.Tag{
		ID:    ulid.Make().String(),
		Name:  tagName,
		Owner: currentUser.ID,
	})
	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong! Try again"))
		return
	}
	w.WriteHeader(http.StatusOK)
	if inserted > 0 {
		if err := h.templates.ExecuteTemplate(w, "tag", map[string]string{
			"Name": tagName,
		}); err != nil {
			logger.Err(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong! Try again"))
			return
		}
	}
}

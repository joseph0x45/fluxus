package handler

import (
	"fluxus/logger"
	"fluxus/models"
	"fmt"
	"net/http"

	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RenderAuthPage(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "auth", models.PageData{Title: "Fluxus | Authenticate"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}
	return string(bytes), nil
}

func passwordMatchesHash(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash), []byte(pwd),
	) == nil
}

func (h *Handler) HandleAuthentication(w http.ResponseWriter, r *http.Request) {
	username, password := r.Form.Get("username"), r.Form.Get("password")
	if len(password) >= 76 {
		w.Write([]byte("Password too long"))
		return
	}
	existingUser, err := h.conn.GetUser("username", username)
	if err != nil {
		logger.Err(err.Error())
		w.Write([]byte("An error occurred. Please try again"))
		return
	}
	if existingUser == nil {
		//create new user
		hash, err := hashPassword(password)
		if err != nil {
			logger.Err(err.Error())
			w.Write([]byte("An error occurred. Please try again"))
			return
		}
		newUser := &models.User{
			ID:       ulid.Make().String(),
			Username: username,
			Password: hash,
		}
		err = h.conn.InsertUser(newUser)
		if err != nil {
			logger.Err(err.Error())
			w.Write([]byte("An error occurred. Please try again"))
		}
		newSession := &models.Session{
			ID: "",
		}
		return
	}
	if !passwordMatchesHash(password, existingUser.Password) {
		w.Write([]byte("This Username is already taken"))
		return
	}
}

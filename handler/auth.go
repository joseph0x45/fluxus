package handler

import (
	"context"
	"fluxus/logger"
	"fluxus/models"
	"fmt"
	"net/http"
	"time"

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
	userID := ""
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("An error occurred. Please try again"))
		return
	}
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
		userID = newUser.ID
	} else {
		if !passwordMatchesHash(password, existingUser.Password) {
			w.Write([]byte("This Username is already taken"))
			return
		}
		userID = existingUser.ID
	}
	newSession := &models.Session{
		ID:          ulid.Make().String(),
		SessionUser: userID,
	}
	if err := h.conn.InsertSession(newSession); err != nil {
		logger.Err(err.Error())
		w.Write([]byte("An error occurred. Please try again"))
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    newSession.ID,
		Path:     "/",
		HttpOnly: h.InReleaseMode(),
		Expires:  time.Now().AddDate(10, 0, 0),
		Secure:   h.InReleaseMode(),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("HX-Redirect", "/")
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			logger.Err("Failed to get cookie: ", err.Error())
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		session, err := h.conn.GetSession(sessionCookie.Value)
		if err != nil {
			logger.Err(err.Error())
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if session == nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		user, err := h.conn.GetUser("id", session.SessionUser)
		if err != nil {
			logger.Err(err.Error())
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if user == nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

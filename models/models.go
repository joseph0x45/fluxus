package models

import "html/template"

type PageData struct {
	StylesCSS template.CSS
	AlpineJS template.JS
	Title     string
}

type User struct {
	ID       string `json:"id" db:"user"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Session struct {
	ID            string `json:"id" db:"id"`
	SessionUserID string `json:"session_user_id" db:"session_user_id"`
}

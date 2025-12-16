package models

type PageData struct {
	Title   string
	Payload map[string]any
}

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Session struct {
	ID          string `json:"id" db:"id"`
	SessionUser string `json:"session_user" db:"session_user"`
}

package models

type PageData struct {
	Title   string
	Payload map[string]any
}

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	SafeMode bool   `json:"safe_mode" db:"safe_mode"`
}

type Session struct {
	ID          string `json:"id" db:"id"`
	SessionUser string `json:"session_user" db:"session_user"`
}

type Account struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Balance int    `json:"balance" db:"balance"`
	Owner   string `json:"owner" db:"owner"`
}

type Tag struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

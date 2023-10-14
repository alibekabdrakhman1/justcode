package model

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

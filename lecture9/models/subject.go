package models

type Subject struct {
	Id       int      `json:"id"`
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}

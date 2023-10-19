package models

type Teacher struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
}

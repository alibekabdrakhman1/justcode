package model

type UserToken struct {
	Id           int    `db:"id"`
	Token        string `db:"token"`
	RefreshToken string `db:"refresh_token"`
	UserId       int    `db:"user_id"`
}

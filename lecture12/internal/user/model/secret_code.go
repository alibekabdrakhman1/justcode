package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type SecretCode struct {
	Id     int
	Code   string
	UserId int
}

type MyJWTClaims struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	*jwt.RegisteredClaims
}
type ContextData struct {
	UserID string
}

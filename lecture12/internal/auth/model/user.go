package model

type Auth struct {
	Login    string
	Password string
}

type JwtUserToken struct {
	Token        string
	RefreshToken string
}

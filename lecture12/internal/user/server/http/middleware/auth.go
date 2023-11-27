package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type JWTAuth struct {
	jwtKey []byte
}

func NewJWTAuth(cfg *config.Config) *JWTAuth {
	return &JWTAuth{jwtKey: []byte(cfg.Auth.JwtSecretKey)}
}

func (m *JWTAuth) ValidateToken(signedToken string) (*model.MyJWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.MyJWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return m.jwtKey, nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.MyJWTClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	return claims, nil
}

func (m *JWTAuth) ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := extractToken(c.Request())

		claims, err := m.ValidateToken(token)
		fmt.Println(claims, err)

		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		ctx := context.WithValue(c.Request().Context(), model.ContextData{}, claims.Id)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

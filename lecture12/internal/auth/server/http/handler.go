package http

import (
	"encoding/json"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/model"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type Manager struct {
	UserToken IUserTokenHandler
}

type UserTokenHandler struct {
	Service *service.Service
}

func NewUserTokenHandler(s *service.Service) *UserTokenHandler {
	return &UserTokenHandler{
		Service: s,
	}
}

func NewManager(srv *service.Service) *Manager {
	return &Manager{NewUserTokenHandler(srv)}
}

func (h *UserTokenHandler) Login(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	request := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	tokenRequest := model.Auth{
		Login:    request.Login,
		Password: request.Password,
	}

	userToken, err := h.Service.UserToken.GenerateToken(c.Request().Context(), tokenRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	response := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	return c.JSON(http.StatusCreated, response)
}

type IUserTokenHandler interface {
	Login(c echo.Context) error
}

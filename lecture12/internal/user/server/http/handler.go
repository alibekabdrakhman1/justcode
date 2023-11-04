package http

import (
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/service"
	"github.com/labstack/echo/v4"
)

type Manager struct {
	User IUserHandler
}

type UserHandler struct {
	Service *service.Service
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(s *service.Service) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func NewManager(srv *service.Service) *Manager {
	return &Manager{NewUserHandler(srv)}
}

type IUserHandler interface {
	GetAllUsers(c echo.Context) error
	GetUserByID(c echo.Context) error
}

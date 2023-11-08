package http

import (
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Manager struct {
	User IUserHandler
}

type UserHandler struct {
	Service *service.Service
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.User.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error getting id from param")
	}
	user, err := h.Service.User.GetUserById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
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

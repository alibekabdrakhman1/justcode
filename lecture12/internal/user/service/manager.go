package service

import (
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/repository"
)

type Service struct {
	User IUserService
}

func NewManager(storage *repository.Repository) (*Service, error) {
	userService := NewUserService(storage)

	return &Service{
		User: userService,
	}, nil
}

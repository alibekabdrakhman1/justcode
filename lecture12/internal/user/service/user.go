package service

import (
	"context"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/model"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/repository"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
}
type UserService struct {
	repository *repository.Repository
}

func NewUserService(r *repository.Repository) *UserService {
	return &UserService{
		repository: r,
	}
}
func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repository.User.GetAllUsers(ctx)
}

func (s *UserService) GetUserById(ctx context.Context, id int) (model.User, error) {
	return s.repository.User.GetUserById(ctx, id)
}

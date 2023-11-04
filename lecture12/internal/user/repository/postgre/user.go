package postgre

import (
	"context"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

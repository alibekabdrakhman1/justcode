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
	var resp []model.User
	err := r.DB.WithContext(ctx).Find(&resp)
	return resp, err.Error
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (model.User, error) {
	var resp model.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).Find(&resp).Error
	return resp, err
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

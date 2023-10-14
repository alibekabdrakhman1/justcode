package postgre

import (
	"github.com/alibekabdrakhman/justcode/lecture8/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAllUsers(c *gin.Context) ([]model.User, error) {
	var resp []model.User
	err := r.DB.WithContext(c.Request.Context()).Find(&resp)
	return resp, err.Error
}

func (r *UserRepository) CreateUser(c *gin.Context, user model.User) (int, error) {
	id := user.ID
	if err := r.DB.WithContext(c.Request.Context()).Create(&user).Error; err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepository) GetUser(c *gin.Context, id int) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(c.Request.Context()).Where("id = ?", id).Find(&res).Error
	return res, err
}

func (r *UserRepository) DeleteUser(c *gin.Context, id int) error {
	return r.DB.WithContext(c.Request.Context()).Delete(&model.User{}, id).Error
}

package storage

import (
	"github.com/alibekabdrakhman/justcode/lecture8/model"
	"github.com/alibekabdrakhman/justcode/lecture8/storage/postgre"
	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	GetAllUsers(c *gin.Context) ([]model.User, error)
	CreateUser(c *gin.Context, user model.User) (int, error)
	GetUser(c *gin.Context, id int) (model.User, error)
	DeleteUser(c *gin.Context, id int) error
}

type Storage struct {
	User IUserRepository
}

func NewStorage(c *gin.Context) (*Storage, error) {
	pgUrl := ""
	DB, err := postgre.Dial(c, pgUrl)
	if err != nil {
		return nil, err
	}
	userRepo := postgre.NewUserRepository(DB)

	storage := Storage{
		User: userRepo,
	}
	return &storage, nil

}

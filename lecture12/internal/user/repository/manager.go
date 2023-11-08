package repository

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/model"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/repository/postgre"
)

func dsn(cfg config.Config) string {
	return fmt.Sprintf("host=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)
}

type Repository struct {
	User IUserRepository
}
type IUserRepository interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	userToken := postgre.NewUserRepository(DB)
	return &Repository{User: userToken}, nil
}

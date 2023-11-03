package repository

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/model"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/repository/postgre"
)

func dsn(cfg config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)
}

type Repository struct {
	UserToken IUserTokenRepository
}
type IUserTokenRepository interface {
	CreateUserToken(ctx context.Context, userToken model.UserToken) error
	UpdateUserToken(ctx context.Context, userToken model.UserToken) error
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	userToken := postgre.NewUserTokenRepository(DB)
	return &Repository{UserToken: userToken}, nil
}

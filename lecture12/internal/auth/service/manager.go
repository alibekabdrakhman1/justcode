package service

import (
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/repository"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/transport"
)

type Service struct {
	UserToken IUserTokenService
}

func NewManager(repo repository.Repository, authConfig config.Auth, userTransport *transport.UserTransport) *Service {
	authService := NewUserTokenService(repo, authConfig, userTransport)
	return &Service{
		UserToken: authService,
	}
}

package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/model"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/repository"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/transport"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserTokenService struct {
	repository        *repository.Repository
	jwtSecretKey      string
	passwordSecretKey string
	userTransport     *transport.UserTransport
}

func NewUserTokenService(repo *repository.Repository, authConfig config.Auth, userTransport *transport.UserTransport) *UserTokenService {
	return &UserTokenService{
		repository:        repo,
		jwtSecretKey:      authConfig.JwtSecretKey,
		passwordSecretKey: authConfig.PasswordSecretKey,
		userTransport:     userTransport,
	}
}

type IUserTokenService interface {
	GenerateToken(ctx context.Context, request model.Auth) (*model.JwtUserToken, error)
}

func (s *UserTokenService) GenerateToken(ctx context.Context, request model.Auth) (*model.JwtUserToken, error) {
	user, err := s.userTransport.GetUser(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}

	generatedPassword := s.generatePassword(request.Password)
	if user.Password != generatedPassword {
		return nil, fmt.Errorf("password is wrong")
	}

	type MyCustomClaims struct {
		UserId int `json:"user_id"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	secretKey := []byte(s.jwtSecretKey)
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := claimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	rClaims := MyCustomClaims{
		user.Id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(40 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	rClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)

	refreshTokenString, err := rClaimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	userToken := model.UserToken{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		UserId:       user.Id,
	}

	err = s.repository.UserToken.CreateUserToken(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("CreateUserToken err: %w", err)
	}

	jwtToken := &model.JwtUserToken{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	return jwtToken, nil
}
func (s *UserTokenService) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(s.passwordSecretKey))
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

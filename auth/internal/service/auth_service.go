package service

import (
	"auth-service/internal/client"
	"auth-service/internal/jwt"
	"errors"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// ✅ Login orqali tokenlar yaratish
func (s *AuthService) Login(username, password string) (*jwt.TokenPair, error) {
	if !client.FindUser(username, password) {
		return nil, errors.New("invalid username or password")
	}

	return jwt.GenerateTokens(username)
}

// ✅ Refresh token orqali yangi token juftligini olish
func (s *AuthService) Refresh(refreshToken string) (*jwt.TokenPair, error) {
	tokens, err := jwt.RefreshTokens(refreshToken)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

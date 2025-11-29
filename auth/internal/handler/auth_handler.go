package handler

import (
	"auth-service/internal/service"
	"context"

	pb "github.com/khbdev/proto-online-test/proto/auth"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: s}
}

// ✅ Login (access + refresh token qaytaradi)
func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokenPair, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		TokenType:    "Bearer",
	}, nil
}

// ✅ RefreshToken — refresh token orqali yangi token juftligini qaytarish
func (h *AuthHandler) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.LoginResponse, error) {
	tokenPair, err := h.svc.Refresh(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		TokenType:    "Bearer",
	}, nil
}

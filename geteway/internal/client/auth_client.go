package client

import (
	"context"
	"fmt"
	"geteway-service/internal/util/connect"
	"log"
	"time"

	authpb "github.com/khbdev/proto-online-test/proto/auth"
	"google.golang.org/grpc"
)

type AuthClient struct {
	conn   *grpc.ClientConn
	client authpb.AuthServiceClient
}

func NewAuthClient() (*AuthClient, error) {
	conn, err := connect.ConnectService("auth-service")
	if err != nil {
		return nil, fmt.Errorf("Auth service bilan ulanish xatosi: %v", err)
	}

	client := authpb.NewAuthServiceClient(conn)

	return &AuthClient{
		client: client,
		conn:   conn,
	}, nil
}

func (a *AuthClient) Close() {
	if a.conn != nil {
		_ = a.conn.Close()
	}
}

// ✅ Login
func (a *AuthClient) Login(username, password string) (*authpb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &authpb.LoginRequest{
		Username: username,
		Password: password,
	}

	res, err := a.client.Login(ctx, req)
	if err != nil {
		log.Printf("[AuthClient] Login xatolik: %v", err)
		return nil, err
	}

	log.Printf("[AuthClient] Login muvaffaqiyatli: %s", username)
	return res, nil
}

// ✅ RefreshToken — refresh token orqali yangi access token olish
func (a *AuthClient) RefreshToken(refreshToken string) (*authpb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &authpb.RefreshRequest{
		RefreshToken: refreshToken,
	}

	res, err := a.client.RefreshToken(ctx, req)
	if err != nil {
		log.Printf("[AuthClient] RefreshToken xatolik: %v", err)
		return nil, err
	}

	log.Printf("[AuthClient] RefreshToken muvaffaqiyatli yangilandi")
	return res, nil
}

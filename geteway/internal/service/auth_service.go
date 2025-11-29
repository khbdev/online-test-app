package service

import (
	"context"
	"fmt"
	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	authpb "github.com/khbdev/proto-online-test/proto/auth"
)

type AuthService struct {
	authClient *client.AuthClient
}

func NewAuthService(authClient *client.AuthClient) *AuthService {
	return &AuthService{authClient: authClient}
}

// ✅ Login
func (s *AuthService) Login(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &authpb.LoginRequest{})
	if err != nil {
		return nil, fmt.Errorf("Login: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*authpb.LoginRequest)
	if !ok {
		return nil, fmt.Errorf("Login: noto‘g‘ri request turi")
	}

	// 🧠 gRPC orqali AuthService'ga so‘rov
	res, err := s.authClient.Login(req.Username, req.Password)
	if err != nil {
		return nil, fmt.Errorf("Login: AuthClient orqali login xatolik: %v", err)
	}

	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("Login: RestGenerate xatolik: %v", err)
	}

	return restRes, nil
}

// ✅ Refresh token orqali yangi access token olish
func (s *AuthService) Refresh(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &authpb.RefreshRequest{})
	if err != nil {
		return nil, fmt.Errorf("Refresh: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*authpb.RefreshRequest)
	if !ok {
		return nil, fmt.Errorf("Refresh: noto‘g‘ri request turi")
	}

	res, err := s.authClient.RefreshToken(req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("Refresh: AuthClient orqali yangilash xatolik: %v", err)
	}

	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("Refresh: RestGenerate xatolik: %v", err)
	}

	return restRes, nil
}

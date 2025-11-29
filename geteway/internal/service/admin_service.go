package service

import (
	"context"

	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	adminpb "github.com/khbdev/proto-online-test/proto/admin"
)

// AdminService — gateway darajasidagi business logika uchun struct
type AdminService struct {
	Client *client.AdminClient
}

// NewAdminService — yangi servis yaratish
func NewAdminService() (*AdminService, error) {
	adminClient, err := client.NewAdminClient()
	if err != nil {
		return nil, fmt.Errorf("admin client yaratishda xato: %v", err)
	}
	return &AdminService{Client: adminClient}, nil
}

// CreateAdmin — yangi admin yaratish
func (s *AdminService) CreateAdmin(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 1️⃣ REST → Proto
	reqMsg, err := adapter.ProtoGenerate(body, &adminpb.CreateAdminRequest{})
	if err != nil {
		return nil, fmt.Errorf("CreateAdmin proto generate error: %v", err)
	}

	req := reqMsg.(*adminpb.CreateAdminRequest)

	// 2️⃣ Client orqali RPC chaqiramiz
	res, err := s.Client.CreateAdmin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("CreateAdmin RPC xato: %v", err)
	}

	// 3️⃣ Proto → REST
	return adapter.RestGenerate(res)
}

// GetAdminList — barcha adminlar
func (s *AdminService) GetAdminList(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.Client.GetAdminList(ctx, &adminpb.GetAdminListRequest{})
	if err != nil {
		return nil, fmt.Errorf("GetAdminList RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

// GetAdminByID — ID orqali admin olish
func (s *AdminService) GetAdminByID(ctx context.Context, id uint64) (map[string]interface{}, error) {
	req := &adminpb.GetAdminByIDRequest{Id: id}
	res, err := s.Client.GetAdminByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetAdminByID RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

// UpdateAdmin — adminni yangilash
func (s *AdminService) UpdateAdmin(ctx context.Context, body []byte) (map[string]interface{}, error) {
	reqMsg, err := adapter.ProtoGenerate(body, &adminpb.UpdateAdminRequest{})
	if err != nil {
		return nil, fmt.Errorf("UpdateAdmin proto generate error: %v", err)
	}

	req := reqMsg.(*adminpb.UpdateAdminRequest)
	res, err := s.Client.UpdateAdmin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UpdateAdmin RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

// DeleteAdmin — adminni o‘chirish
func (s *AdminService) DeleteAdmin(ctx context.Context, id uint64) (map[string]interface{}, error) {
	req := &adminpb.DeleteAdminRequest{Id: id}
	res, err := s.Client.DeleteAdmin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("DeleteAdmin RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

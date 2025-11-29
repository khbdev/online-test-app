package service

import (
	"context"
	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	userpb "github.com/khbdev/proto-online-test/proto/user"
)

// UserService — gateway darajasidagi business logika uchun struct
type UserService struct {
	Client *client.UserClient
}

// NewUserService — yangi servis yaratish
func NewUserService() (*UserService, error) {
	userClient, err := client.NewUserClient()
	if err != nil {
		return nil, fmt.Errorf("user client yaratishda xato: %v", err)
	}
	return &UserService{Client: userClient}, nil
}

// CreateUser — yangi user yaratish
func (s *UserService) CreateUser(ctx context.Context, body []byte) (map[string]interface{}, error) {
	reqMsg, err := adapter.ProtoGenerate(body, &userpb.CreateUserRequest{})
	if err != nil {
		return nil, fmt.Errorf("CreateUser proto generate error: %v", err)
	}
	req := reqMsg.(*userpb.CreateUserRequest)

	res, err := s.Client.CreateUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("CreateUser RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

// GetAllUsers — barcha userlar
func (s *UserService) GetAllUsers(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.Client.GetAllUsers(ctx, &userpb.GetAllUsersRequest{})
	if err != nil {
		return nil, fmt.Errorf("GetAllUsers RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

// GetUserByID — ID orqali user olish
func (s *UserService) GetUserByID(ctx context.Context, id uint64) (map[string]interface{}, error) {
	req := &userpb.GetUserByIDRequest{Id: id}
	res, err := s.Client.GetUserByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

// UpdateUser — user ma’lumotlarini yangilash
func (s *UserService) UpdateUser(ctx context.Context, body []byte) (map[string]interface{}, error) {
	reqMsg, err := adapter.ProtoGenerate(body, &userpb.UpdateUserRequest{})
	if err != nil {
		return nil, fmt.Errorf("UpdateUser proto generate error: %v", err)
	}
	req := reqMsg.(*userpb.UpdateUserRequest)

	res, err := s.Client.UpdateUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

// DeleteUser — userni o‘chirish
func (s *UserService) DeleteUser(ctx context.Context, id uint64) (map[string]interface{}, error) {
	req := &userpb.DeleteUserRequest{Id: id}
	res, err := s.Client.DeleteUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("DeleteUser RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

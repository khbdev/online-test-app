package client

import (
	"context"
	"fmt"
	"time"

	"geteway-service/internal/util/connect"

	userpb "github.com/khbdev/proto-online-test/proto/user"
	"google.golang.org/grpc"
)

// UserClient - User service bilan RPC aloqasi uchun
type UserClient struct {
	conn   *grpc.ClientConn
	client userpb.UserServiceClient
}

// NewUserClient - User service bilan ulanish va client yaratish
func NewUserClient() (*UserClient, error) {
	conn, err := connect.ConnectService("user-service")
	if err != nil {
		return nil, fmt.Errorf("User service bilan ulanish xatosi: %v", err)
	}

	client := userpb.NewUserServiceClient(conn)
	return &UserClient{
		conn:   conn,
		client: client,
	}, nil
}

// CreateUser - yangi user yaratish
func (u *UserClient) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return u.client.CreateUser(ctx, req)
}

// GetAllUsers - barcha userlarni olish
func (u *UserClient) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return u.client.GetAllUsers(ctx, req)
}

// GetUserByID - ID orqali user olish
func (u *UserClient) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return u.client.GetUserByID(ctx, req)
}

// UpdateUser - user ma’lumotlarini yangilash
func (u *UserClient) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return u.client.UpdateUser(ctx, req)
}

// DeleteUser - userni o‘chirish
func (u *UserClient) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return u.client.DeleteUser(ctx, req)
}

// Close - connection yopish
func (u *UserClient) Close() {
	if u.conn != nil {
		_ = u.conn.Close()
	}
}

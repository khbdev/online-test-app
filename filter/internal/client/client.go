package client

import (
	"context"
	"log"
	"time"

userpb "github.com/khbdev/proto-online-test/proto/user"

	"google.golang.org/grpc"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatalf("❌ User service bilan ulanib bo‘lmadi: %v", err)
	}

	log.Println("✅ User Service bilan gRPC ulanish o‘rnatildi!")
	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (u *UserClient) GetAllUsers(ctx context.Context) ([]*userpb.User, error) {
	resp, err := u.client.GetAllUsers(ctx, &userpb.GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

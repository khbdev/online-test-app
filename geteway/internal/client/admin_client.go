package client

import (
	"context"
	"fmt"
	"time"

	"geteway-service/internal/util/connect"

	adminpb "github.com/khbdev/proto-online-test/proto/admin"
	"google.golang.org/grpc"
)


type AdminClient struct {
	conn   *grpc.ClientConn
	client adminpb.AdminServiceClient
}


func NewAdminClient() (*AdminClient, error) {
	conn, err := connect.ConnectService("admin-service")
	if err != nil {
		return nil, fmt.Errorf("Admin service bilan ulanish xatosi: %v", err)
	}
	client := adminpb.NewAdminServiceClient(conn)

		return &AdminClient{
		client: client,
		conn: conn,
	}, nil
}

// CreateAdmin - yangi admin yaratish
func (a *AdminClient) CreateAdmin(ctx context.Context, req *adminpb.CreateAdminRequest) (*adminpb.CreateAdminResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return a.client.CreateAdmin(ctx, req)
}

// GetAdminList - barcha adminlar
func (a *AdminClient) GetAdminList(ctx context.Context, req *adminpb.GetAdminListRequest) (*adminpb.GetAdminListResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return a.client.GetAdminList(ctx, req)
}

// GetAdminByID - id orqali admin olish
func (a *AdminClient) GetAdminByID(ctx context.Context, req *adminpb.GetAdminByIDRequest) (*adminpb.GetAdminByIDResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return a.client.GetAdminByID(ctx, req)
}

// UpdateAdmin - adminni yangilash
func (a *AdminClient) UpdateAdmin(ctx context.Context, req *adminpb.UpdateAdminRequest) (*adminpb.UpdateAdminResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return a.client.UpdateAdmin(ctx, req)
}

// DeleteAdmin - adminni o‘chirish
func (a *AdminClient) DeleteAdmin(ctx context.Context, req *adminpb.DeleteAdminRequest) (*adminpb.DeleteAdminResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return a.client.DeleteAdmin(ctx, req)
}

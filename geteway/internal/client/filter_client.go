package client

import (
	"context"
	"fmt"
	"geteway-service/internal/util/connect"
	filterpb "github.com/khbdev/proto-online-test/proto/filter"
	"google.golang.org/grpc"
	"log"
	"time"
)

type FilterClient struct {
	conn   *grpc.ClientConn
	client filterpb.UserServiceClient
}

// ✅ filter-service bilan ulanish
func NewFilterClient() (*FilterClient, error) {
	conn, err := connect.ConnectService("filter-service")
	if err != nil {
		return nil, fmt.Errorf("filter-service bilan ulanish xatosi: %v", err)
	}

	client := filterpb.NewUserServiceClient(conn)

	return &FilterClient{
		conn:   conn,
		client: client,
	}, nil
}

// ✅ ulanishni yopish
func (f *FilterClient) Close() {
	if f.conn != nil {
		_ = f.conn.Close()
	}
}

// ✅ GetUsers RPC chaqiruv
func (f *FilterClient) GetUsers(req *filterpb.FilterRequest) (*filterpb.UserList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := f.client.GetUsers(ctx, req)
	if err != nil {
		log.Printf("[FilterClient] GetUsers RPC xato: %v", err)
		return nil, err
	}

	log.Printf("[FilterClient] %d ta foydalanuvchi topildi", len(res.Users))
	return res, nil
}

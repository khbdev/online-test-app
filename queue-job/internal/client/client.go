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
	conn   *grpc.ClientConn
}

// NewUserClient - 50053 portga ulanadi
func NewUserClient() *UserClient {
	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ gRPC ulanish xatolik: %v", err)
	}

	c := userpb.NewUserServiceClient(conn)
	return &UserClient{client: c, conn: conn}
}

// Close - connection yopadi
func (u *UserClient) Close() {
	u.conn.Close()
}

// CreateUser - user-service dagi CreateUser RPC’ni chaqiradi
func (u *UserClient) CreateUser(req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ✅ Real ma'lumotlar bilan RPC chaqiriladi
	res, err := u.client.CreateUser(ctx, &userpb.CreateUserRequest{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Phone:           req.Phone,
		Email:           req.Email,
		TgUsername:      req.TgUsername,
		Bolimlar:        req.Bolimlar,
		Savollar:        req.Savollar,
		Javoblar:        req.Javoblar,
		TogriJavoblar:   req.TogriJavoblar,
		NatogriJavoblar: req.NatogriJavoblar,
		ScorePercent:    req.ScorePercent,
		Description:     req.Description,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

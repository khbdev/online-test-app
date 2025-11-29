package client

import (
	"auth-service/internal/config"
	"context"
	"log"
	"time"

	pb "github.com/khbdev/proto-online-test/proto/admin"
	"google.golang.org/grpc"
)

func FindUser(username, password string) bool {
	conn, err := grpc.Dial(config.GRPC.Address(), grpc.WithInsecure())
	if err != nil {
		log.Printf("❌ Failed to connect to admin-service: %v", err)
		return false
	}
	defer conn.Close()

	client := pb.NewAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.VerifyAdmin(ctx, &pb.VerifyAdminRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Printf("❌ VerifyAdmin error: %v", err)
		return false
	}

	return resp.Valid
}

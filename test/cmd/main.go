package main

import (
	"context"
	"log"
	"time"

	"github.com/khbdev/proto-online-test/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. gRPC serverga ulanamiz
	conn, err := grpc.Dial("auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	// 2. Request tayyorlaymiz
	req := &auth.LoginRequest{
		Username: "admin",
		Password: "123456",
	}

	// 3. Context bilan timeout belgilaymiz
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Login RPC chaqiramiz
	resp, err := client.Login(ctx, req)
	if err != nil {
		log.Fatalf("⚠️ Login RPC error: %v", err)
	}

	// 5. Javobni chiqaramiz
	log.Println("✅ Access Token:", resp.AccessToken)
	log.Println("✅ Refresh Token:", resp.RefreshToken)
	log.Println("🔑 Token Type:", resp.TokenType)
	log.Println("⏳ Expires In:", resp.ExpiresIn)
}
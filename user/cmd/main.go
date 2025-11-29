package main

import (
	"fmt"
	"log"
	"net"

	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/repostory"
	"user-service/internal/service"

	"github.com/joho/godotenv"
	userpb "github.com/khbdev/proto-online-test/proto/user"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Initialize database connection
	config.InitDB()

	// Get gRPC port from environment
	portPtr := config.GetEnv("GRPC_PORT")
	if portPtr == nil {
		log.Fatal("Missing GRPC_PORT in .env file")
	}
	port := *portPtr

	// Initialize layers: repository → service → handler
	repo := repostory.NewUserRepository(config.DB)
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc)

	// Start TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create and register gRPC server
	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, userHandler)

	fmt.Printf(" User gRPC server is running on %s\n", port)

	// Serve gRPC
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

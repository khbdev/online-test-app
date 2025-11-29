package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/service"
	"log"
	"net"

	pb "github.com/khbdev/proto-online-test/proto/auth"

	"google.golang.org/grpc"
)

func main(){
	// env config 
	config.InitJWT()
	config.InitPort()
	config.InitGRPC()


lis, err := net.Listen("tcp", config.Server.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authService := service.NewAuthService()
	pb.RegisterAuthServiceServer(s, handler.NewAuthHandler(authService))
	log.Println("auth grpc server running on: 50051")
		if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
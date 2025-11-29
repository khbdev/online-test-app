package main

import (
	"admin-service/internal/config"
	"admin-service/internal/handler"
	repository "admin-service/internal/repostory"
	"admin-service/internal/service"
	"fmt"
	"log"
	"net"

	pb "github.com/khbdev/proto-online-test/proto/admin"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }

    if err := config.InitDB(); err != nil {
        log.Fatal(err)
    }

	config.InitPort()

    log.Println("Database connected successfully")

	repo := repository.NewAdminRepository(config.DB)
    
	service := service.NewAdminService(repo)

   handler := handler.NewAdminHandler(service)

   lis, err := net.Listen("tcp", config.Server.GRPCPort)
	if err != nil {
		log.Fatal("Port band:", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdminServiceServer(s, handler)


	fmt.Println("Admin gRPC server is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Server Error:", err)
	}


}

package main

import (
	"filter-service/internal/client"
	"filter-service/internal/handler"
	"filter-service/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
		filterpb "github.com/khbdev/proto-online-test/proto/filter"
)


func main(){
  client := 	client.NewUserClient("user:50053")

	service := service.NewFilterService(client)

	handler := handler.NewFilterHandler(service)


	lis, err := net.Listen("tcp", ":50057") 
	if err != nil {
		log.Fatalf("Port band: %v", err)
	}
		grpcServer := grpc.NewServer()
	filterpb.RegisterUserServiceServer(grpcServer, handler)

	log.Println("Test Generation Service running on port :50057")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
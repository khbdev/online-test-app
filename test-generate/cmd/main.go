package main

import (
	"log"
	"net"

	"test-generation-servis/internal/client"
	"test-generation-servis/internal/config"
	"test-generation-servis/internal/handler"
	repo "test-generation-servis/internal/repostory"
	"test-generation-servis/internal/service"

	"github.com/joho/godotenv"
	pb "github.com/khbdev/proto-online-test/proto/generate"
	"google.golang.org/grpc"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Failed to load .env file: %v", err)
    }

    // Redis ni ishga tushirish
    config.InitRedis() // <<< shu chaqirilishi shart

    if config.RedisClient == nil {
        log.Fatal("RedisClient is nil")
    }

	// Repository yaratish
	testRepo := repo.NewRepository(config.RedisClient)

	// Section gRPC client
	sectionClient := client.NewSectionClient(config.GetEnv("TEST_SECTION_SERVICE"))

	// TestService yaratish
	testService := service.NewTestService(testRepo, sectionClient, config.GetEnv("TEST_GENERATE_URL"))

	// Handler yaratish
	testHandler := handler.NewTestHandler(testService)

	// GRPC_PORT ni .env dan olish
	port := config.GetEnv("GRPS_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Port band: %v", err)
	}

	// gRPC server ishga tushirish
	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, testHandler)

	log.Printf("Test Generation Service running on port: %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}

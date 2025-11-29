package main

import (
	"log"
	"net"
	"queue-job-service/internal/config"
	"queue-job-service/internal/handler"
	"queue-job-service/internal/service"

	jobpb "github.com/khbdev/proto-online-test/proto/job"

	"google.golang.org/grpc"
)

func main() {
	// =========================
	// 🔹 Configlar
	// =========================
	rmq := config.NewRabbitMQ()
	defer rmq.Conn.Close()
	defer rmq.Channel.Close()

	config.InitRedis()

	// =========================
	// 🔹 Job Consumer (background)
	// =========================
	go service.ConsumeJobs()

	// =========================
	// 🔹 GRPC Server ishga tushirish
	// =========================
	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("❌ GRPC listen xato: %v", err)
	}

	grpcServer := grpc.NewServer()
	jobpb.RegisterJobServiceServer(grpcServer, &handler.JobServiceServer{})

	log.Println("🚀 GRPC server ishlamoqda: :50056")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ GRPC server xato: %v", err)
	}
}

package main

import (
	"log"
	"net"

	"test-section-service/internal/config"
	"test-section-service/internal/handler"
	repositorys "test-section-service/internal/repostorys"
	"test-section-service/internal/service"

	testpb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/grpc"
)

func main() {

	config.InitDB()


	sectionRepo := repositorys.NewSectionRepository(config.DB)
	questionRepo := repositorys.NewQuestionRepository(config.DB)
	optionRepo := repositorys.NewOptionRepository(config.DB)


	sectionService := service.NewSectionService(sectionRepo)
	questionService := service.NewQuestionService(questionRepo)
	optionService := service.NewOptionService(optionRepo)


	testHandler := handler.NewTestHandler(sectionService, questionService, optionService)


	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal("Port band:", err)
	}

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, testHandler)

	log.Println("Test Section Service running on port :50054")

	if err := s.Serve(lis); err != nil {
		log.Fatal(" Serverda xatolik:", err)
	}
}

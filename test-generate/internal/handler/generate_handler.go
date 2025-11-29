package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"test-generation-servis/internal/service"

	pb "github.com/khbdev/proto-online-test/proto/generate"
)

type TestHandler struct {
	pb.UnimplementedTestServiceServer
	service *service.TestService
}

func NewTestHandler(svc *service.TestService) *TestHandler {
	return &TestHandler{service: svc}
}

// --- Generate Test ---
func (h *TestHandler) GenerateTest(ctx context.Context, req *pb.GenerateTestRequest) (*pb.GenerateTestResponse, error) {
    // int32 -> uint64
    sectionIDs := make([]uint64, len(req.SectionIds))
    for i, id := range req.SectionIds {
        sectionIDs[i] = uint64(id)
    }

    link, err := h.service.GenerateTest(req.Name, sectionIDs)
    if err != nil {
        log.Println("Failed to generate test:", err)
        return nil, err
    }

    return &pb.GenerateTestResponse{Link: link}, nil
}

// --- Get Test by Key ---
func (h *TestHandler) GetTest(ctx context.Context, req *pb.GetTestRequest) (*pb.GetTestResponse, error) {
	data, err := h.service.GetTest(req.Key)
	if err != nil {
		log.Println("Failed to get test:", err)
		return nil, err
	}

	// TestData ni JSON formatga o‘tkazish
	testJSON, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal test data: %w", err)
	}

	return &pb.GetTestResponse{
		TestJson: string(testJSON),
	}, nil
}

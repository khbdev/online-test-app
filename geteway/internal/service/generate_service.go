package service

import (
	"context"
	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	generatepb "github.com/khbdev/proto-online-test/proto/generate"
)

type GenerateService struct {
	generateClient *client.GenerateClient
}

// ✅ yangi servis yaratish
func NewGenerateService(generateClient *client.GenerateClient) *GenerateService {
	return &GenerateService{generateClient: generateClient}
}

// ✅ Test yaratish (GenerateTest)
func (s *GenerateService) GenerateTest(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 🔹 REST → Proto
	protoReq, err := adapter.ProtoGenerate(body, &generatepb.GenerateTestRequest{})
	if err != nil {
		return nil, fmt.Errorf("GenerateTest: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*generatepb.GenerateTestRequest)
	if !ok {
		return nil, fmt.Errorf("GenerateTest: noto‘g‘ri request turi")
	}

	// 🔹 gRPC RPC chaqirish
	res, err := s.generateClient.GenerateTest(req.Name, req.SectionIds)
	if err != nil {
		return nil, fmt.Errorf("GenerateTest: RPC xato: %v", err)
	}

	// 🔹 Proto → REST
	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("GenerateTest: RestGenerate xato: %v", err)
	}

	return restRes, nil
}

// ✅ Testni olish (GetTest)
func (s *GenerateService) GetTest(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 🔹 REST → Proto
	protoReq, err := adapter.ProtoGenerate(body, &generatepb.GetTestRequest{})
	if err != nil {
		return nil, fmt.Errorf("GetTest: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*generatepb.GetTestRequest)
	if !ok {
		return nil, fmt.Errorf("GetTest: noto‘g‘ri request turi")
	}

	// 🔹 gRPC RPC chaqirish
	res, err := s.generateClient.GetTest(req.Key)
	if err != nil {
		return nil, fmt.Errorf("GetTest: RPC xato: %v", err)
	}

	// 🔹 Proto → REST
	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("GetTest: RestGenerate xato: %v", err)
	}

	
	return restRes, nil
}

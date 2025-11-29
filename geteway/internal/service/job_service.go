package service

import (
	"context"
	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"
	jobpb "github.com/khbdev/proto-online-test/proto/job"
)

type JobService struct {
	jobClient *client.JobClient
}

// ✅ yangi servis yaratish
func NewJobService(jobClient *client.JobClient) *JobService {
	return &JobService{jobClient: jobClient}
}

// ✅ Testni topshirish (SubmitTest)
func (s *JobService) SubmitTest(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 🔹 REST → Proto
	protoReq, err := adapter.ProtoGenerate(body, &jobpb.SubmitTestRequest{})
	if err != nil {
		return nil, fmt.Errorf("SubmitTest: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*jobpb.SubmitTestRequest)
	if !ok {
		return nil, fmt.Errorf("SubmitTest: noto‘g‘ri request turi")
	}

	// 🔹 gRPC RPC chaqirish
	res, err := s.jobClient.SubmitTest(req)
	if err != nil {
		return nil, fmt.Errorf("SubmitTest: RPC xato: %v", err)
	}

	// 🔹 Proto → REST
	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("SubmitTest: RestGenerate xato: %v", err)
	}

	return restRes, nil
}

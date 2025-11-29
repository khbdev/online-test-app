package service

import (
	"context"
	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	exampb "github.com/khbdev/proto-online-test/proto/exam"
)

// ExamService — gateway darajasidagi business logika uchun struct
type ExamService struct {
	Client *client.ExamClient
}

// NewExamService — yangi servis yaratish
func NewExamService() (*ExamService, error) {
	examClient, err := client.NewExamClient()
	if err != nil {
		return nil, fmt.Errorf("exam client yaratishda xato: %v", err)
	}
	return &ExamService{Client: examClient}, nil
}

// CreateExam — yangi exam yaratish
func (s *ExamService) CreateExam(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 1️⃣ REST → Proto
	reqMsg, err := adapter.ProtoGenerate(body, &exampb.CreateExamRequest{})
	if err != nil {
		return nil, fmt.Errorf("CreateExam proto generate error: %v", err)
	}

	req := reqMsg.(*exampb.CreateExamRequest)

	// 2️⃣ Client orqali RPC chaqiramiz
	res, err := s.Client.CreateExam(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("CreateExam RPC xato: %v", err)
	}

	// 3️⃣ Proto → REST
	return adapter.RestGenerate(res)
}

// GetExamList — barcha examlar
func (s *ExamService) GetExamList(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.Client.GetExamList(ctx, &exampb.GetExamListRequest{})
	if err != nil {
		return nil, fmt.Errorf("GetExamList RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

// GetExamByID — ID orqali exam olish
func (s *ExamService) GetExamByID(ctx context.Context, id uint64) (map[string]interface{}, error) {
	req := &exampb.GetExamByIDRequest{Id: id}
	res, err := s.Client.GetExamByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetExamByID RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

// UpdateExam — examni yangilash
func (s *ExamService) UpdateExam(ctx context.Context, body []byte) (map[string]interface{}, error) {
	reqMsg, err := adapter.ProtoGenerate(body, &exampb.UpdateExamRequest{})
	if err != nil {
		return nil, fmt.Errorf("UpdateExam proto generate error: %v", err)
	}

	req := reqMsg.(*exampb.UpdateExamRequest)
	res, err := s.Client.UpdateExam(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UpdateExam RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

// DeleteExam — examni o‘chirish
func (s *ExamService) DeleteExam(ctx context.Context, id uint64) (map[string]interface{}, error) {
	req := &exampb.DeleteExamRequest{Id: id}
	res, err := s.Client.DeleteExam(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("DeleteExam RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

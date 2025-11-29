package client

import (
	"context"
	"fmt"
	"geteway-service/internal/util/connect"
	"time"

	exampb "github.com/khbdev/proto-online-test/proto/exam"

	"google.golang.org/grpc"
)

type ExamClient struct {
	conn   *grpc.ClientConn
	client exampb.ExamServiceClient
}

func NewExamClient() (*ExamClient, error) {
	conn, err := connect.ConnectService("exam-service")
	if err != nil {
		return nil, fmt.Errorf("Exam service bilan ulanish xatosi: %v", err)
	}
	client := exampb.NewExamServiceClient(conn)

	return &ExamClient{
		client: client,
		conn:   conn,
	}, nil
}

// CreateExam - yangi exam yaratish
func (c *ExamClient) CreateExam(ctx context.Context, req *exampb.CreateExamRequest) (*exampb.CreateExamResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.client.CreateExam(ctx, req)
}

// GetExamList - barcha examlar
func (c *ExamClient) GetExamList(ctx context.Context, req *exampb.GetExamListRequest) (*exampb.GetExamListResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.client.GetExamList(ctx, req)
}

// GetExamByID - id orqali exam olish
func (c *ExamClient) GetExamByID(ctx context.Context, req *exampb.GetExamByIDRequest) (*exampb.GetExamByIDResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.client.GetExamByID(ctx, req)
}

// UpdateExam - examni yangilash
func (c *ExamClient) UpdateExam(ctx context.Context, req *exampb.UpdateExamRequest) (*exampb.UpdateExamResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.client.UpdateExam(ctx, req)
}

// DeleteExam - examni o‘chirish
func (c *ExamClient) DeleteExam(ctx context.Context, req *exampb.DeleteExamRequest) (*exampb.DeleteExamResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.client.DeleteExam(ctx, req)
}

package client

import (
	"context"
	"fmt"
	"geteway-service/internal/util/connect"
	jobpb "github.com/khbdev/proto-online-test/proto/job"
	"google.golang.org/grpc"
	"log"
	"time"
)

type JobClient struct {
	conn   *grpc.ClientConn
	client jobpb.JobServiceClient
}

// ✅ job-service bilan ulanish
func NewJobClient() (*JobClient, error) {
	conn, err := connect.ConnectService("job-service")
	if err != nil {
		return nil, fmt.Errorf("job-service bilan ulanish xatosi: %v", err)
	}

	client := jobpb.NewJobServiceClient(conn)

	return &JobClient{
		client: client,
		conn:   conn,
	}, nil
}

// ✅ ulanishni yopish
func (j *JobClient) Close() {
	if j.conn != nil {
		_ = j.conn.Close()
	}
}

// ✅ SubmitTest RPC chaqiruv
func (j *JobClient) SubmitTest(req *jobpb.SubmitTestRequest) (*jobpb.SubmitTestResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := j.client.SubmitTest(ctx, req)
	if err != nil {
		log.Printf("[JobClient] SubmitTest RPC xato: %v", err)
		return nil, err
	}

	log.Printf("[JobClient] Test yuborildi: key=%s", res.Key)
	return res, nil
}

package client

import (
	"context"
	"fmt"
	"geteway-service/internal/util/connect"
	"log"
	"time"

	generatepb "github.com/khbdev/proto-online-test/proto/generate"
	"google.golang.org/grpc"
)

type GenerateClient struct {
	conn   *grpc.ClientConn
	client generatepb.TestServiceClient
}

// ✅ generate-service bilan ulanish
func NewGenerateClient() (*GenerateClient, error) {
	conn, err := connect.ConnectService("generate-service")
	if err != nil {
		return nil, fmt.Errorf("generate-service bilan ulanish xatosi: %v", err)
	}

	client := generatepb.NewTestServiceClient(conn)

	return &GenerateClient{
		client: client,
		conn:   conn,
	}, nil
}


func (g *GenerateClient) Close() {
	if g.conn != nil {
		_ = g.conn.Close()
	}
}


func (g *GenerateClient) GenerateTest(name string, sectionIDs []int32) (*generatepb.GenerateTestResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &generatepb.GenerateTestRequest{
		Name:       name,
		SectionIds: sectionIDs,
	}

	res, err := g.client.GenerateTest(ctx, req)
	if err != nil {
		log.Printf("[GenerateClient] GenerateTest RPC xato: %v", err)
		return nil, err
	}

	log.Printf("[GenerateClient] Test yaratildi: %s", res.Link)
	return res, nil
}


func (g *GenerateClient) GetTest(key string) (*generatepb.GetTestResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &generatepb.GetTestRequest{
		Key: key,
	}

	res, err := g.client.GetTest(ctx, req)
	if err != nil {
		log.Printf("[GenerateClient] GetTest RPC xato: %v", err)
		return nil, err
	}

	log.Printf("[GenerateClient] Test olindi: key=%s", key)
	return res, nil
}

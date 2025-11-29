package client

import (
	"context"
	"fmt"
	"geteway-service/internal/util/connect"
	"log"
	"time"

	testpb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TestClient struct {
	conn   *grpc.ClientConn
	client testpb.TestServiceClient
}

// ✅ test-service bilan ulanish
func NewTestClient() (*TestClient, error) {
	conn, err := connect.ConnectService("test-service")
	if err != nil {
		return nil, fmt.Errorf("test-service bilan ulanish xatosi: %v", err)
	}

	client := testpb.NewTestServiceClient(conn)
	return &TestClient{
		client: client,
		conn:   conn,
	}, nil
}

// ✅ ulanishni yopish
func (t *TestClient) Close() {
	if t.conn != nil {
		_ = t.conn.Close()
	}
}

//
// ========== SECTION CRUD ==========
//

func (t *TestClient) CreateSection(req *testpb.CreateSectionRequest) (*testpb.Section, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.CreateSection(ctx, req)
	if err != nil {
		log.Printf("[TestClient] CreateSection RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) GetSectionById(id uint64) (*testpb.GetSectionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.GetSectionById(ctx, &testpb.SectionID{Id: id})
	if err != nil {
		log.Printf("[TestClient] GetSectionById RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) GetAllSections() (*testpb.GetAllSectionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.GetAllSections(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("[TestClient] GetAllSections RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) UpdateSection(req *testpb.UpdateSectionRequest) (*testpb.Section, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.UpdateSection(ctx, req)
	if err != nil {
		log.Printf("[TestClient] UpdateSection RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) DeleteSection(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteSection(ctx, &testpb.SectionID{Id: id})
	if err != nil {
		log.Printf("[TestClient] DeleteSection RPC xato: %v", err)
		return err
	}
	return nil
}

//
// ========== QUESTION CRUD ==========
//

func (t *TestClient) CreateQuestion(req *testpb.CreateQuestionRequest) (*testpb.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.CreateQuestion(ctx, req)
	if err != nil {
		log.Printf("[TestClient] CreateQuestion RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) GetQuestionById(id uint64) (*testpb.GetQuestionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.GetQuestionById(ctx, &testpb.QuestionID{Id: id})
	if err != nil {
		log.Printf("[TestClient] GetQuestionById RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) GetAllQuestions() (*testpb.GetAllQuestionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.GetAllQuestions(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("[TestClient] GetAllQuestions RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) UpdateQuestion(req *testpb.UpdateQuestionRequest) (*testpb.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.UpdateQuestion(ctx, req)
	if err != nil {
		log.Printf("[TestClient] UpdateQuestion RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) DeleteQuestion(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteQuestion(ctx, &testpb.QuestionID{Id: id})
	if err != nil {
		log.Printf("[TestClient] DeleteQuestion RPC xato: %v", err)
		return err
	}
	return nil
}

//
// ========== OPTION CRUD ==========
//

func (t *TestClient) CreateOption(req *testpb.CreateOptionRequest) (*testpb.Option, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.CreateOption(ctx, req)
	if err != nil {
		log.Printf("[TestClient] CreateOption RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) GetOptionById(id uint64) (*testpb.GetOptionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.GetOptionById(ctx, &testpb.OptionID{Id: id})
	if err != nil {
		log.Printf("[TestClient] GetOptionById RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) GetAllOptions() (*testpb.GetAllOptionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.GetAllOptions(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("[TestClient] GetAllOptions RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) UpdateOption(req *testpb.UpdateOptionRequest) (*testpb.Option, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := t.client.UpdateOption(ctx, req)
	if err != nil {
		log.Printf("[TestClient] UpdateOption RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *TestClient) DeleteOption(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteOption(ctx, &testpb.OptionID{Id: id})
	if err != nil {
		log.Printf("[TestClient] DeleteOption RPC xato: %v", err)
		return err
	}
	return nil
}

//
// ========== FULL STRUCTURE ==========
//

func (t *TestClient) GetFullSectionStructure(sectionID uint64) (*testpb.GetFullSectionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &testpb.GetFullSectionRequest{SectionId: sectionID}
	res, err := t.client.GetFullSectionStructure(ctx, req)
	if err != nil {
		log.Printf("[TestClient] GetFullSectionStructure RPC xato: %v", err)
		return nil, err
	}
	return res, nil
}

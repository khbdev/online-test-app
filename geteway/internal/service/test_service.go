package service

import (
	"context"
	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	testpb "github.com/khbdev/proto-online-test/proto/test"
)

// ✅ TestService — test bilan ishlovchi service
type TestService struct {
	testClient *client.TestClient
}

// ✅ Yangi TestService yaratish
func NewTestService(testClient *client.TestClient) *TestService {
	return &TestService{testClient: testClient}
}

//
// ========== SECTION CRUD ==========
//

func (s *TestService) CreateSection(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &testpb.CreateSectionRequest{})
	if err != nil {
		return nil, fmt.Errorf("CreateSection: ProtoGenerate xatolik: %v", err)
	}
	req := protoReq.(*testpb.CreateSectionRequest)

	res, err := s.testClient.CreateSection(req)
	if err != nil {
		return nil, fmt.Errorf("CreateSection RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

func (s *TestService) GetSectionById(ctx context.Context, id uint64) (map[string]interface{}, error) {
	res, err := s.testClient.GetSectionById(id)
	if err != nil {
		return nil, fmt.Errorf("GetSectionById RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

func (s *TestService) GetAllSections(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.testClient.GetAllSections()
	if err != nil {
		return nil, fmt.Errorf("GetAllSections RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

func (s *TestService) UpdateSection(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &testpb.UpdateSectionRequest{})
	if err != nil {
		return nil, fmt.Errorf("UpdateSection: ProtoGenerate xatolik: %v", err)
	}
	req := protoReq.(*testpb.UpdateSectionRequest)

	res, err := s.testClient.UpdateSection(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateSection RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

func (s *TestService) DeleteSection(ctx context.Context, id uint64) (map[string]interface{}, error) {
	err := s.testClient.DeleteSection(id)
	if err != nil {
		return nil, fmt.Errorf("DeleteSection RPC xato: %v", err)
	}
	return map[string]interface{}{"success": true, "message": "Section o‘chirildi"}, nil
}

//
// ========== QUESTION CRUD ==========
//

func (s *TestService) CreateQuestion(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &testpb.CreateQuestionRequest{})
	if err != nil {
		return nil, fmt.Errorf("CreateQuestion: ProtoGenerate xatolik: %v", err)
	}
	req := protoReq.(*testpb.CreateQuestionRequest)

	res, err := s.testClient.CreateQuestion(req)
	if err != nil {
		return nil, fmt.Errorf("CreateQuestion RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

func (s *TestService) GetQuestionById(ctx context.Context, id uint64) (map[string]interface{}, error) {
	res, err := s.testClient.GetQuestionById(id)
	if err != nil {
		return nil, fmt.Errorf("GetQuestionById RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

func (s *TestService) GetAllQuestions(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.testClient.GetAllQuestions()
	if err != nil {
		return nil, fmt.Errorf("GetAllQuestions RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

func (s *TestService) UpdateQuestion(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &testpb.UpdateQuestionRequest{})
	if err != nil {
		return nil, fmt.Errorf("UpdateQuestion: ProtoGenerate xatolik: %v", err)
	}
	req := protoReq.(*testpb.UpdateQuestionRequest)

	res, err := s.testClient.UpdateQuestion(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateQuestion RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

func (s *TestService) DeleteQuestion(ctx context.Context, id uint64) (map[string]interface{}, error) {
	err := s.testClient.DeleteQuestion(id)
	if err != nil {
		return nil, fmt.Errorf("DeleteQuestion RPC xato: %v", err)
	}
	return map[string]interface{}{"success": true, "message": "Question o‘chirildi"}, nil
}

//
// ========== OPTION CRUD ==========
//

func (s *TestService) CreateOption(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &testpb.CreateOptionRequest{})
	if err != nil {
		return nil, fmt.Errorf("CreateOption: ProtoGenerate xatolik: %v", err)
	}
	req := protoReq.(*testpb.CreateOptionRequest)

	res, err := s.testClient.CreateOption(req)
	if err != nil {
		return nil, fmt.Errorf("CreateOption RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

func (s *TestService) GetOptionById(ctx context.Context, id uint64) (map[string]interface{}, error) {
	res, err := s.testClient.GetOptionById(id)
	if err != nil {
		return nil, fmt.Errorf("GetOptionById RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

func (s *TestService) GetAllOptions(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.testClient.GetAllOptions()
	if err != nil {
		return nil, fmt.Errorf("GetAllOptions RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

func (s *TestService) UpdateOption(ctx context.Context, body []byte) (map[string]interface{}, error) {
	protoReq, err := adapter.ProtoGenerate(body, &testpb.UpdateOptionRequest{})
	if err != nil {
		return nil, fmt.Errorf("UpdateOption: ProtoGenerate xatolik: %v", err)
	}
	req := protoReq.(*testpb.UpdateOptionRequest)

	res, err := s.testClient.UpdateOption(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateOption RPC xato: %v", err)
	}

	return adapter.RestGenerate(res)
}

func (s *TestService) DeleteOption(ctx context.Context, id uint64) (map[string]interface{}, error) {
	err := s.testClient.DeleteOption(id)
	if err != nil {
		return nil, fmt.Errorf("DeleteOption RPC xato: %v", err)
	}
	return map[string]interface{}{"success": true, "message": "Option o‘chirildi"}, nil
}

//
// ========== FULL STRUCTURE ==========
//

func (s *TestService) GetFullSectionStructure(ctx context.Context, sectionID uint64) (map[string]interface{}, error) {
	res, err := s.testClient.GetFullSectionStructure(sectionID)
	if err != nil {
		return nil, fmt.Errorf("GetFullSectionStructure RPC xato: %v", err)
	}
	return adapter.RestGenerate(res)
}

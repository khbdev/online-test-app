package service

import (
	"context"
	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"
	filterpb "github.com/khbdev/proto-online-test/proto/filter"
)

type FilterService struct {
	filterClient *client.FilterClient
}

func NewFilterService(filterClient *client.FilterClient) *FilterService {
	return &FilterService{filterClient: filterClient}
}

func (s *FilterService) GetUsers(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 🔹 REST → Proto
	protoReq, err := adapter.ProtoGenerate(body, &filterpb.FilterRequest{})
	if err != nil {
		return nil, fmt.Errorf("GetUsers: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*filterpb.FilterRequest)
	if !ok {
		return nil, fmt.Errorf("GetUsers: noto‘g‘ri request turi")
	}

	// 🔹 gRPC RPC chaqirish
	res, err := s.filterClient.GetUsers(req)
	if err != nil {
		return nil, fmt.Errorf("GetUsers: RPC xato: %v", err)
	}

	// 🔹 Proto → REST
	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("GetUsers: RestGenerate xato: %v", err)
	}

	return restRes, nil
}

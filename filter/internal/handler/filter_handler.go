package handler

import (
	"context"
	"encoding/json"
	"filter-service/internal/service"
	"log"

	filterpb "github.com/khbdev/proto-online-test/proto/filter"
	userpb "github.com/khbdev/proto-online-test/proto/user"
)

type FilterHandler struct {
	FilterService *service.FilterService
	filterpb.UnimplementedUserServiceServer
}

func NewFilterHandler(filterService *service.FilterService) *FilterHandler {
	return &FilterHandler{
		FilterService: filterService,
	}
}

func (h *FilterHandler) GetUsers(ctx context.Context, req *filterpb.FilterRequest) (*filterpb.UserList, error) {

	users, err := h.FilterService.FilterUsers(ctx, req.Year, req.FirstName, req.LastName, req.Phone)
	if err != nil {
		log.Printf("Filter qilishda xatolik: %v", err)
		return nil, err
	}

	var pbUsers []*filterpb.FilterUser
	for _, u := range users {
		pbUsers = append(pbUsers, mapUser(u))
	}

	return &filterpb.UserList{Users: pbUsers}, nil
}

// ✅ userpb.User ichida JSON bo'lgan maydonlarni decode qilamiz
func mapUser(u *userpb.User) *filterpb.FilterUser {
	var bolimlar []*filterpb.Section
	var savollar []*filterpb.Question
	var javoblar []*filterpb.Answer

	// JSON parse qilish
	_ = json.Unmarshal([]byte(u.Bolimlar), &bolimlar)
	_ = json.Unmarshal([]byte(u.Savollar), &savollar)
	_ = json.Unmarshal([]byte(u.Javoblar), &javoblar)

	return &filterpb.FilterUser{
		Id:              int64(u.Id),
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
		DeletedAt:       "",
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Phone:           u.Phone,
		Email:           u.Email,
		TgUsername:      u.TgUsername,
		Bolimlar:        bolimlar,
		Savollar:        savollar,
		Javoblar:        javoblar,
		TogriJavoblar:   u.TogriJavoblar,
		NatogriJavoblar: u.NatogriJavoblar,
		ScorePercent:    float64(u.ScorePercent),
		Description:     u.Description,
	}
}

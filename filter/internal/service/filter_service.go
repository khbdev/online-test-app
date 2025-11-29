package service

import (
	"context"
	"filter-service/internal/client"
	"log"
	"strings"
	"time"

	userpb "github.com/khbdev/proto-online-test/proto/user"
)

type FilterService struct {
	UserClient *client.UserClient
}

func NewFilterService(userClient *client.UserClient) *FilterService {
	return &FilterService{UserClient: userClient}
}

func (f *FilterService) FilterUsers(ctx context.Context, year string, firstName string, lastName string, phone string) ([]*userpb.User, error) {
	
	users, err := f.UserClient.GetAllUsers(ctx)
	if err != nil {
		log.Printf(" Userlarni olishda xatolik: %v", err)
		return nil, err
	}

	var filtered []*userpb.User

	for _, u := range users {
		match := true

	
		if year != "" {
			t, err := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
			if err == nil {
				if t.Format("2006") != year {
					match = false
				}
			}
		}


		if firstName != "" && !strings.Contains(strings.ToLower(u.FirstName), strings.ToLower(firstName)) {
			match = false
		}


		if lastName != "" && !strings.Contains(strings.ToLower(u.LastName), strings.ToLower(lastName)) {
			match = false
		}

		if phone != "" && !strings.Contains(u.Phone, phone) {
			match = false
		}

		if match {
			filtered = append(filtered, u)
		}
	}

	return filtered, nil
}

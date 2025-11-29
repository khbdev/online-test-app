package service

import (
	"fmt"
	"time"

	client "test-generation-servis/internal/client"
	repository "test-generation-servis/internal/repostory"

	testpb "github.com/khbdev/proto-online-test/proto/test"
)



type TestService struct {
	Repo          *repository.TestRepository 
	SectionClient *client.SectionClient    
	Domain        string                   
}

func NewTestService(repo *repository.TestRepository, sc *client.SectionClient, domain string) *TestService {
	return &TestService{
		Repo:          repo,
		SectionClient: sc,
		Domain:        domain,
	}
}

// POST: GenerateTest
func (s *TestService) GenerateTest(name string, sectionIDs []uint64) (string, error) {
	var allSections []*testpb.Section

	for _, id := range sectionIDs {
		section, err := s.SectionClient.GetSection(id)
		if err != nil {
			return "", fmt.Errorf("section %d olishda xatolik: %w", id, err)
		}

		
		allSections = append(allSections, section)
	}

	key := fmt.Sprintf("%s_%d", name, time.Now().UnixNano())

	data := repository.TestData{
		Name:     name,
		Sections: allSections,
	}

	if err := s.Repo.Set(name, key, data, 30*time.Minute); err != nil {
		return "", fmt.Errorf("Redisga saqlashda xatolik: %w", err)
	}

	link := fmt.Sprintf("%s/test/%s", s.Domain, key)
	return link, nil
}


func (s *TestService) GetTest(key string) (*repository.TestData, error) {
	data, err := s.Repo.Get(key)
	if err != nil {
		return nil, fmt.Errorf("testni Redisdan olishda xatolik: %w", err)
	}


	clientSections := make([]*testpb.Section, len(data.Sections))
	for i, sec := range data.Sections {
		secCopy := *sec // copy
		for _, q := range secCopy.Questions {
			for _, o := range q.Options {
				o.IsCorrect = false
			}
		}
		clientSections[i] = &secCopy
	}

	return &repository.TestData{
		Name:     data.Name,
		Sections: clientSections,
	}, nil
}

package service

import (
	"test-section-service/internal/models"
	repositorys "test-section-service/internal/repostorys"
)

type SectionService struct {
	Repo *repositorys.SectionRepository
}

func NewSectionService(repo *repositorys.SectionRepository) *SectionService {
	return &SectionService{Repo: repo}
}

func (s *SectionService) CreateSection(section *models.Section) error {
	return s.Repo.Create(section)
}

func (s *SectionService) GetAllSections() ([]models.Section, error) {
	return s.Repo.GetAll()
}

func (s *SectionService) GetSectionByID(id uint) (*models.Section, error) {
	return s.Repo.GetByID(id)
}

func (s *SectionService) UpdateSection(section *models.Section) error {
	return s.Repo.Update(section)
}

func (s *SectionService) DeleteSection(id uint) error {
	return s.Repo.Delete(id)
}

// 🔥 Queue Job uchun kerakli to‘liq section structure
func (s *SectionService) GetFullSectionStructure(sectionID uint) (*models.Section, error) {
	return s.Repo.GetFullSectionStructure(sectionID)
}

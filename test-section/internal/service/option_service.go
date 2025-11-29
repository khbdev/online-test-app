package service

import (
	"test-section-service/internal/models"
	repositorys "test-section-service/internal/repostorys"
)

type OptionService struct {
	Repo *repositorys.OptionRepository
}

func NewOptionService(repo *repositorys.OptionRepository) *OptionService {
	return &OptionService{Repo: repo}
}

func (s *OptionService) CreateOption(o *models.Option) error {
	return s.Repo.Create(o)
}

func (s *OptionService) GetAllOptions() ([]models.Option, error) {
	return s.Repo.GetAll()
}

func (s *OptionService) GetOptionByID(id uint) (*models.Option, error) {
	return s.Repo.GetByID(id)
}

func (s *OptionService) UpdateOption(o *models.Option) error {
	return s.Repo.Update(o)
}

func (s *OptionService) DeleteOption(id uint) error {
	return s.Repo.Delete(id)
}

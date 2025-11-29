package service

import (
	"test-section-service/internal/models"
	repositorys "test-section-service/internal/repostorys"
)

type QuestionService struct {
	repo *repositorys.QuestionRepository
}

func NewQuestionService(repo *repositorys.QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) Create(question *models.Question) (*models.Question, error) {
	if err := s.repo.Create(question); err != nil {
		return nil, err
	}
	return question, nil
}

func (s *QuestionService) GetByID(id uint) (*models.Question, error) {
	return s.repo.GetByID(id)
}

func (s *QuestionService) GetAll() ([]models.Question, error) {
	return s.repo.GetAll()
}

func (s *QuestionService) Update(question *models.Question) (*models.Question, error) {
	if err := s.repo.Update(question); err != nil {
		return nil, err
	}
	return question, nil
}

func (s *QuestionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

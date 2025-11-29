package repositorys

import (
	"test-section-service/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{DB: db}
}

func (r *QuestionRepository) Create(q *models.Question) error {
	return r.DB.Create(q).Error
}

func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question
	err := r.DB.Preload("Options").Find(&questions).Error
	return questions, err
}

func (r *QuestionRepository) GetByID(id uint) (*models.Question, error) {
	var q models.Question
	err := r.DB.Preload("Options").First(&q, id).Error
	return &q, err
}

func (r *QuestionRepository) Update(q *models.Question) error {
	return r.DB.Model(&models.Question{}).
		Where("id = ?", q.ID).
		Updates(map[string]interface{}{
			"section_id": q.SectionID,
			"text":       q.Text,
		}).Error
}

func (r *QuestionRepository) Delete(id uint) error {
	return r.DB.Unscoped().Delete(&models.Question{}, id).Error
}

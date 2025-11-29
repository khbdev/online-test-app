package repositorys

import (
	"test-section-service/internal/models"

	"gorm.io/gorm"
)

type OptionRepository struct {
	DB *gorm.DB
}

func NewOptionRepository(db *gorm.DB) *OptionRepository {
	return &OptionRepository{DB: db}
}

func (r *OptionRepository) Create(o *models.Option) error {
	return r.DB.Create(o).Error
}

func (r *OptionRepository) GetAll() ([]models.Option, error) {
	var options []models.Option
	err := r.DB.Find(&options).Error
	return options, err
}

func (r *OptionRepository) GetByID(id uint) (*models.Option, error) {
	var o models.Option
	err := r.DB.First(&o, id).Error
	return &o, err
}

func (r *OptionRepository) Update(o *models.Option) error {
	return r.DB.Model(&models.Option{}).
		Where("id = ?", o.ID).
		Updates(map[string]interface{}{
			"text":       o.Text,
			"is_correct": o.IsCorrect,
		}).Error
}

func (r *OptionRepository) Delete(id uint) error {
	return r.DB.Unscoped().Delete(&models.Option{}, id).Error
}

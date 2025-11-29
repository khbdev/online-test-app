package repositorys

import (
	"test-section-service/internal/models"

	"gorm.io/gorm"
)

type SectionRepository struct {
	DB *gorm.DB
}


func NewSectionRepository(db *gorm.DB) *SectionRepository {
	return &SectionRepository{DB: db}
}


func (r *SectionRepository) Create(section *models.Section) error {
	return r.DB.Create(section).Error
}


func (r *SectionRepository) GetAll() ([]models.Section, error) {
	var sections []models.Section
	err := r.DB.Preload("Questions.Options").Find(&sections).Error
	return sections, err
}


func (r *SectionRepository) GetByID(id uint) (*models.Section, error) {
	var section models.Section
	err := r.DB.Preload("Questions.Options").First(&section, id).Error
	return &section, err
}

func (r *SectionRepository) Update(section *models.Section) error {
	return r.DB.Save(section).Error
}


func (r *SectionRepository) Delete(id uint) error {
	return r.DB.Unscoped().Delete(&models.Section{}, id).Error
}

func (r *SectionRepository) GetFullSectionStructure(sectionID uint) (*models.Section, error) {
	var section models.Section
	err := r.DB.Preload("Questions.Options").First(&section, sectionID).Error
	if err != nil {
		return nil, err
	}
	return &section, nil
}
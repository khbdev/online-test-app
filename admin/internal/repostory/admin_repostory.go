package repository

import (
	"admin-service/internal/model"
	"errors"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) GetAll() ([]model.Admin, error) {
	var admins []model.Admin
	if err := r.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *AdminRepository) GetByID(id uint) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.First(&admin, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) Create(username, password string) (*model.Admin, error) {
	admin := &model.Admin{
		Username: username,
		Password: password, 
	}
	if err := r.db.Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepository) Update(id uint, username, password string) (*model.Admin, error) {
	admin, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, gorm.ErrRecordNotFound
	}

	admin.Username = username
	admin.Password = password 

	if err := r.db.Save(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepository) Delete(id uint) error {
	admin, err := r.GetByID(id)
	if err != nil {
		return err
	}
	if admin == nil {
		return gorm.ErrRecordNotFound
	}
	if err := r.db.Delete(admin).Error; err != nil {
		return err
	}
	return nil
}

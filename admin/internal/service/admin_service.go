package service

import (
	"admin-service/internal/model"
	repository "admin-service/internal/repostory"

	"admin-service/internal/util"
	"errors"
)

type AdminService struct {
	repo *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}


func (s *AdminService) Create(username, password string) (*model.Admin, error) {
	hashed, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(username, hashed)
}


func (s *AdminService) Update(id uint, username, password string) (*model.Admin, error) {
	hashed, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(id, username, hashed)
}


func (s *AdminService) GetAll() ([]model.Admin, error) {
	return s.repo.GetAll()
}


func (s *AdminService) GetByID(id uint) (*model.Admin, error) {
	return s.repo.GetByID(id)
}


func (s *AdminService) Delete(id uint) error {
	return s.repo.Delete(id)
}


func (s *AdminService) VerifyAdmin(username, password string) (bool, error) {
	admins, err := s.repo.GetAll()
	if err != nil {
		return false, err
	}

	for _, admin := range admins {
		if admin.Username == username {
			if util.CheckPassword(password, admin.Password) {
				return true, nil
			}
			return false, errors.New("invalid password")
		}
	}

	return false, errors.New("admin not found")
}

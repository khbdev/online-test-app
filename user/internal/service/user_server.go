package service

import (
	"encoding/json"
	models "user-service/internal/model"
	"user-service/internal/repostory"

)

type UserService struct {
	repo *repostory.UserRepository
}

func NewUserService(repo *repostory.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create yangi user
func (s *UserService) Create(
	firstName, lastName, phone, email, tgUsername,
	bolimlar, savollar, javoblar, description string,
	togriJavoblar, natogriJavoblar, scorePercent int,
) (*models.User, error) {

	user := &models.User{
		FirstName:       firstName,
		LastName:        lastName,
		Phone:           phone,
		Email:           email,
		TgUsername:      tgUsername,
		Bolimlar:        bolimlar,
		Savollar:        savollar,
		Javoblar:        javoblar,
		TogriJavoblar:   togriJavoblar,
		NatogriJavoblar: natogriJavoblar,
		ScorePercent:    scorePercent,
		Description:     description,
	}

	u, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return parseUserJSONFields(u), nil
}

// GetAll barcha userlar
func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range users {
	 parseUserJSONFields(&users[i])
	}

	return users, nil
}

// GetByID ID orqali user
func (s *UserService) GetByID(id uint) (*models.User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return parseUserJSONFields(u), nil
}

// Update user
func (s *UserService) Update(
	id uint,
	firstName, lastName, phone, email, tgUsername,
	bolimlar, savollar, javoblar, description string,
	togriJavoblar, natogriJavoblar, scorePercent int,
) (*models.User, error) {

	updatedUser := &models.User{
		FirstName:       firstName,
		LastName:        lastName,
		Phone:           phone,
		Email:           email,
		TgUsername:      tgUsername,
		Bolimlar:        bolimlar,
		Savollar:        savollar,
		Javoblar:        javoblar,
		TogriJavoblar:   togriJavoblar,
		NatogriJavoblar: natogriJavoblar,
		ScorePercent:    scorePercent,
		Description:     description,
	}

	u, err := s.repo.Update(id, updatedUser)
	if err != nil {
		return nil, err
	}

	return parseUserJSONFields(u), nil
}

// Delete user
func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// parseUserJSONFields - string JSON maydonlarni parse qilib obyektga aylantiradi
func parseUserJSONFields(u *models.User) *models.User {
	if u == nil {
		return nil
	}

	var parsed interface{}

	if u.Bolimlar != "" {
		if err := json.Unmarshal([]byte(u.Bolimlar), &parsed); err == nil {
			u.BolimlarParsed = parsed
		}
	}

	if u.Savollar != "" {
		if err := json.Unmarshal([]byte(u.Savollar), &parsed); err == nil {
			u.SavollarParsed = parsed
		}
	}

	if u.Javoblar != "" {
		if err := json.Unmarshal([]byte(u.Javoblar), &parsed); err == nil {
			u.JavoblarParsed = parsed
		}
	}

	return u
}

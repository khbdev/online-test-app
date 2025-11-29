package repostory

import (
	"errors"
	models "user-service/internal/model"

	"gorm.io/gorm"
)



type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Select("id, first_name, last_name, phone, email, tg_username, togri_javoblar, natogri_javoblar, score_percent, description").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}


func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func (r *UserRepository) Update(id uint, updatedUser *models.User) (*models.User, error) {
	user, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gorm.ErrRecordNotFound
	}


	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Phone = updatedUser.Phone
	user.Email = updatedUser.Email
	user.TgUsername = updatedUser.TgUsername
	user.Bolimlar = updatedUser.Bolimlar
	user.Savollar = updatedUser.Savollar
	user.Javoblar = updatedUser.Javoblar
	user.TogriJavoblar = updatedUser.TogriJavoblar
	user.NatogriJavoblar = updatedUser.NatogriJavoblar
	user.ScorePercent = updatedUser.ScorePercent
	user.Description = updatedUser.Description
	

	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func (r *UserRepository) Delete(id uint) error {
	user, err := r.GetByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return gorm.ErrRecordNotFound
	}
	if err := r.db.Unscoped().Delete(user).Error; err != nil {
		return err
	}
	return nil
}
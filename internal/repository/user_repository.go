package repository

import (
	"github.com/mr-emerald-wolf/brew-backend/internal/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{db: database.DB}
}

func (ur UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur UserRepository) Create(user domain.User) (*domain.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepository) Find(uuid string) (*domain.User, error) {
	var user domain.User
	err := ur.db.Where("uuid = ?", uuid).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepository) Update(uuid string, update domain.User) (*domain.User, error) {
	var user domain.User
	err := ur.db.Where("uuid = ?", uuid).First(&user).Error

	if err != nil {
		return nil, err
	}

	// Update user fields with the new data
	user.Name = update.Name
	user.Phone = update.Phone
	user.Email = update.Email

	err = ur.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepository) Delete(uuid string) error {
	var user domain.User
	err := ur.db.Where("uuid = ?", uuid).First(&user).Error

	if err != nil {
		return err
	}

	err = ur.db.Delete(&user).Unscoped().Error
	if err != nil {
		return err
	}

	return nil
}

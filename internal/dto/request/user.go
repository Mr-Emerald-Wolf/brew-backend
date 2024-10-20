package dto

import (
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
)

type UserCreateRequest struct {
	Name     string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"full_name"`
	Phone    string `json:"phone" validate:"numeric"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

func (u UserCreateRequest) ToDomain() db.User {
	return db.User{
		Name:         u.Name,
		Phone:        u.Phone,
		Email:        u.Email,
		PasswordHash: u.Password,
	}
}

func (u UserUpdateRequest) ToDomain() db.User {
	return db.User{
		Name:         u.Name,
		Phone:        u.Phone,
		Email:        u.Email,
		PasswordHash: u.Password,
	}
}

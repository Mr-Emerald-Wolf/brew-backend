package dto

import (
	"github.com/google/uuid"
	"github.com/mr-emerald-wolf/brew-backend/internal/domain"
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

func (u UserCreateRequest) ToDomain() domain.User {
	return domain.User{
		UUID:     uuid.New(),
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u UserUpdateRequest) ToDomain() domain.User {
	return domain.User{
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Password: u.Password,
	}
}

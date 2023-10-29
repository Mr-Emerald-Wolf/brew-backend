package dto

import (
	"github.com/google/uuid"
	"github.com/mr-emerald-wolf/brew-backend/internal/domain"
)

type UserCreateRequest struct {
	Name  string `json:"full_name" validate:"required"`
	Phone string `json:"phone" validate:"required,numeric"`
	Email string `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	Name  string `json:"full_name" validate:"required"`
	Phone string `json:"phone" validate:"required,numeric"`
	Email string `json:"email" validate:"required,email"`
}

func (u UserCreateRequest) ToDomain() domain.User {
	return domain.User{
		UUID:  uuid.New(),
		Name:  u.Name,
		Phone: u.Phone,
		Email: u.Email,
	}
}

func (u UserUpdateRequest) ToDomain() domain.User {
	return domain.User{
		Name:  u.Name,
		Phone: u.Phone,
		Email: u.Email,
	}
}

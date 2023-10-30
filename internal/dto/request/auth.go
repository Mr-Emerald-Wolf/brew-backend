package dto

import "github.com/mr-emerald-wolf/brew-backend/internal/domain"

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (u AuthRequest) ToDomain() domain.User {
	return domain.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

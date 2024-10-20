package dto

import (
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (u AuthRequest) ToDomain() db.User {
	return db.User{
		Email:        u.Email,
		PasswordHash: u.Password,
	}
}

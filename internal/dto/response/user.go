package dto

import (
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
)

type UserResponse struct {
	Name      string `json:"full_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToUserDTO(u db.User) UserResponse {
	return UserResponse{
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt.Time.String(),
		UpdatedAt: u.UpdatedAt.Time.String(),
	}
}

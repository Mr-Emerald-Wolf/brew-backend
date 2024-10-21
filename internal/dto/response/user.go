package dto

import (
	"fmt"

	"github.com/mr-emerald-wolf/brew-backend/internal/db"
)

type UserResponse struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"full_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToUserDTO(u db.User) UserResponse {
	return UserResponse{
		Uuid:      fmt.Sprintf("%x", u.Uuid.Bytes),
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt.Time.String(),
		UpdatedAt: u.UpdatedAt.Time.String(),
	}
}

package domain

import (
	"strconv"

	"github.com/google/uuid"
	dto "github.com/mr-emerald-wolf/brew-backend/internal/dto/response"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name  string    `gorm:"type:varchar(255);not null"`
	Email string    `gorm:"type:varchar(255);not null"`
	Phone string    `gorm:"type:varchar(255);not null"`
}

func (u User) ToDTO() dto.UserResponse {
	return dto.UserResponse{
		ID:        strconv.FormatUint(uint64(u.ID), 10),
		UUID:      u.UUID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt.GoString(),
		UpdatedAt: u.UpdatedAt.GoString(),
	}
}

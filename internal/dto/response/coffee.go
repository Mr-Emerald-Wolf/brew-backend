package dto

import (
	"fmt"

	"github.com/mr-emerald-wolf/brew-backend/internal/db"
)

type CoffeeResponse struct {
	ID        int32  `json:"coffee_id"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Origin    string `json:"origin"`
	Roast     string `json:"roast"`
	Process   string `json:"process"`
	Price     int32  `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToCoffeeDTO(c db.Coffee) CoffeeResponse {
	return CoffeeResponse{
		ID:        c.ID,
		UUID:      fmt.Sprintf("%x", c.Uuid.Bytes),
		Name:      c.Name,
		Origin:    c.Origin,
		Roast:     c.Roast,
		Process:   c.Process,
		Price:     c.Price,
		CreatedAt: c.CreatedAt.Time.String(),
		UpdatedAt: c.UpdatedAt.Time.String(),
	}
}

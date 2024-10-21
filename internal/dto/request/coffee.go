package dto

import (
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
)

type CoffeeCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Origin  string `json:"origin" validate:"required"`
	Roast   string `json:"roast" validate:"required"`
	Process string `json:"process" validate:"required"`
	Price   int32  `json:"price" validate:"required"`
}

type CoffeeUpdateRequest struct {
	Name    string `json:"name"`
	Origin  string `json:"origin"`
	Roast   string `json:"roast"`
	Process string `json:"process"`
	Price   int32  `json:"price"`
}

func (c CoffeeCreateRequest) ToDomain() db.Coffee {
	return db.Coffee{
		Name:    c.Name,
		Origin:  c.Origin,
		Roast:   c.Roast,
		Process: c.Process,
		Price:   c.Price,
	}
}

func (c CoffeeUpdateRequest) ToDomain() db.Coffee {
	return db.Coffee{
		Name:    c.Name,
		Origin:  c.Origin,
		Roast:   c.Roast,
		Process: c.Process,
		Price:   c.Price,
	}
}

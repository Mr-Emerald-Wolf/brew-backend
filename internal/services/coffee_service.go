package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	res "github.com/mr-emerald-wolf/brew-backend/internal/dto/response"
)

type ICoffeeService interface {
	CreateCoffee(req.CoffeeCreateRequest, int32) (*res.CoffeeResponse, error)
	FindAllCoffees() ([]res.CoffeeResponse, error)
	FindCoffee(string) (*res.CoffeeResponse, error)
	UpdateCoffee(string, req.CoffeeUpdateRequest) (*res.CoffeeResponse, error)
	DeleteCoffee(string) error
}

type CoffeeService struct {
	repo *db.Queries
}

func NewCoffeeService(repo *db.Queries) *CoffeeService {
	return &CoffeeService{repo: repo}
}

func (cs *CoffeeService) CreateCoffee(r req.CoffeeCreateRequest, userID int32) (*res.CoffeeResponse, error) {
	newCoffee, err := cs.repo.CreateCoffee(context.Background(), db.CreateCoffeeParams{
		UserID:  userID,
		Name:    r.Name,
		Origin:  r.Origin,
		Roast:   r.Roast,
		Process: r.Process,
		Price:   r.Price,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create coffee: %w", err)
	}

	// Convert the db.Coffee object to a response DTO
	response := res.ToCoffeeDTO(newCoffee)

	return &response, nil
}

func (cs *CoffeeService) FindAllCoffees() ([]res.CoffeeResponse, error) {
	allCoffees, err := cs.repo.GetAllCoffees(context.Background())
	if err != nil {
		return nil, err
	}

	response := make([]res.CoffeeResponse, 0)
	for _, coffee := range allCoffees {
		response = append(response, res.ToCoffeeDTO(coffee))
	}

	return response, nil
}

func (cs *CoffeeService) FindCoffee(uuidStr string) (*res.CoffeeResponse, error) {
	// Parse the UUID string
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("parse UUID: %w", err)
	}

	// Convert to pgtype.UUID
	pgUUID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	findCoffee, err := cs.repo.GetCoffeeByUUID(context.Background(), pgUUID)
	if err != nil {
		return nil, err
	}

	response := res.ToCoffeeDTO(findCoffee)

	return &response, nil
}

func (cs *CoffeeService) UpdateCoffee(uuidStr string, r req.CoffeeUpdateRequest) (*res.CoffeeResponse, error) {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("parse UUID: %w", err)
	}
	pgUUID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}
	updatedCoffee, err := cs.repo.UpdateCoffee(context.Background(), db.UpdateCoffeeParams{
		Uuid:    pgUUID,
		Name:    r.Name,
		Origin:  r.Origin,
		Roast:   r.Roast,
		Process: r.Process,
		Price:   r.Price,
	})
	if err != nil {
		return nil, fmt.Errorf("update coffee: %w", err)
	}
	response := res.ToCoffeeDTO(updatedCoffee)
	return &response, nil
}

func (cs *CoffeeService) DeleteCoffee(uuidStr string) error {
	// Parse the UUID string
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return fmt.Errorf("failed to parse UUID: %w", err)
	}

	// Convert to pgtype.UUID
	UUID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}
	err = cs.repo.DeleteCoffee(context.Background(), UUID)
	return err
}

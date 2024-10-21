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

type IUserService interface {
	CreateUser(req.UserCreateRequest) (*res.UserResponse, error)
	FindAllUsers() ([]res.UserResponse, error)
	FindUser(string) (*res.UserResponse, error)
	UpdateUser(string, req.UserUpdateRequest) (*res.UserResponse, error)
	DeleteUser(string) error
}

type UserService struct {
	repo *db.Queries
}

func NewUserService(repo *db.Queries) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(r req.UserCreateRequest) (*res.UserResponse, error) {
	hashedPassword, err := hashPassword(r.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user, err := us.repo.CreateUser(context.Background(), db.CreateUserParams{
		Name:         r.Name,
		Email:        r.Email,
		PasswordHash: hashedPassword,
		Phone:        r.Phone,
	})
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	response := res.ToUserDTO(user)
	return &response, nil
}

func (us *UserService) FindUser(uuidStr string) (*res.UserResponse, error) {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("parse UUID: %w", err)
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	user, err := us.repo.GetUserByUUID(context.Background(), pgUUID)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	response := res.ToUserDTO(user)
	return &response, nil
}

func (us *UserService) FindAllUsers() ([]res.UserResponse, error) {
	users, err := us.repo.GetAllUsers(context.Background())
	if err != nil {
		return nil, fmt.Errorf("find all users: %w", err)
	}

	response := make([]res.UserResponse, len(users))
	for i, user := range users {
		response[i] = res.ToUserDTO(user)
	}

	return response, nil
}

func (us *UserService) UpdateUser(uuidStr string, r req.UserUpdateRequest) (*res.UserResponse, error) {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("parse UUID: %w", err)
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	user, err := us.repo.UpdateUser(context.Background(), db.UpdateUserParams{
		Uuid:  pgUUID,
		Name:  r.Name,
		Email: r.Email,
		Phone: r.Phone,
	})
	if err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	response := res.ToUserDTO(user)
	return &response, nil
}

func (us *UserService) DeleteUser(uuidStr string) error {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return fmt.Errorf("parse UUID: %w", err)
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	if err := us.repo.DeleteUser(context.Background(), pgUUID); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

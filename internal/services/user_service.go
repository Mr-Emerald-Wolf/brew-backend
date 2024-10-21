package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	res "github.com/mr-emerald-wolf/brew-backend/internal/dto/response"
)

type IUserService interface {
	CreateUser(req.UserCreateRequest) (*res.UserResponse, error)
	FindAllUsers() ([]res.UserResponse, error)
	FindUser(pgtype.UUID) (*res.UserResponse, error)
	UpdateUser(pgtype.UUID, req.UserUpdateRequest) (*res.UserResponse, error)
	DeleteUser(pgtype.UUID) error
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

func (us *UserService) FindUser(uuid pgtype.UUID) (*res.UserResponse, error) {

	user, err := us.repo.GetUserByUUID(context.Background(), uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no user found")
		}
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

func (us *UserService) UpdateUser(uuid pgtype.UUID, r req.UserUpdateRequest) (*res.UserResponse, error) {

	user, err := us.repo.UpdateUser(context.Background(), db.UpdateUserParams{
		Uuid:  uuid,
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

func (us *UserService) DeleteUser(uuid pgtype.UUID) error {

	if err := us.repo.DeleteUser(context.Background(), uuid); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

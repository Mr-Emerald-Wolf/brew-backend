package services

import (
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	res "github.com/mr-emerald-wolf/brew-backend/internal/dto/response"

	"github.com/mr-emerald-wolf/brew-backend/internal/repository"
)

type IUserService interface {
	CreateUser(req.UserCreateRequest) (*res.UserResponse, error)
	FindAllUsers() (*[]res.UserResponse, error)
	FindUser(uuid string) (*res.UserResponse, error)
}

type UserService struct {
	r repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{r: repo}
}

func (us UserService) CreateUser(u req.UserCreateRequest) (*res.UserResponse, error) {
	user := u.ToDomain()
	newUser, err := us.r.Create(user)
	if err != nil {
		return nil, err
	}
	response := newUser.ToDTO()

	return &response, nil
}

func (us UserService) FindUser(uuid string) (*res.UserResponse, error) {
	findUser, err := us.r.Find(uuid)
	if err != nil {
		return nil, err
	}
	response := findUser.ToDTO()

	return &response, nil
}

func (us UserService) FindAllUsers() (*[]res.UserResponse, error) {
	allUsers, err := us.r.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]res.UserResponse, 0)
	for _, c := range allUsers {
		response = append(response, c.ToDTO())
	}

	return &response, nil
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
	"github.com/mr-emerald-wolf/brew-backend/internal/utils"
)

type IUserHandler interface {
	NewCustomer(*fiber.Ctx) error
	GetAllCustomers(*fiber.Ctx) error
	GetCustomer(*fiber.Ctx) error
}

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(us services.IUserService) UserHandler {
	return UserHandler{
		service: us,
	}
}

func (uh *UserHandler) NewUser(c *fiber.Ctx) error {
	var payload req.UserCreateRequest
	err := c.BodyParser(&payload)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	response, err := uh.service.CreateUser(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	c.Status(fiber.StatusCreated).JSON(response)
	return nil
}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {

	users, err := uh.service.FindAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {

	uuid := c.Params("uuid")

	user, err := uh.service.FindUser(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

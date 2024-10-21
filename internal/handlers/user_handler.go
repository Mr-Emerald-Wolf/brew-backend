package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
	"github.com/mr-emerald-wolf/brew-backend/internal/utils"
)

type IUserHandler interface {
	NewCustomer(*fiber.Ctx) error
	GetAllCustomers(*fiber.Ctx) error
	GetCustomer(*fiber.Ctx) error
	UpdateUser(*fiber.Ctx) error
	DeleteUser(*fiber.Ctx) error
	Me(*fiber.Ctx) error
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": errors})
	}

	response, err := uh.service.CreateUser(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": true, "user": response})
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

	user_uuid := c.Params("uuid")
	parsedUUID, err := uuid.Parse(user_uuid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	user, err := uh.service.FindUser(pgUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	uuid := c.Locals("user").(pgtype.UUID)

	var payload req.UserUpdateRequest
	err := c.BodyParser(&payload)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	response, err := uh.service.UpdateUser(uuid, payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {

	uuid := c.Locals("user").(pgtype.UUID)

	err := uh.service.DeleteUser(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": "User Deleted"})
}

func (uh *UserHandler) Me(c *fiber.Ctx) error {

	uuid := c.Locals("user").(pgtype.UUID)

	user, err := uh.service.FindUser(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

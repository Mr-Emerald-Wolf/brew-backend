package handlers

import (
	"github.com/gofiber/fiber/v2"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
	"github.com/mr-emerald-wolf/brew-backend/internal/utils"
)

type IAuthHandler interface {
	Login(*fiber.Ctx) error
	Logout(*fiber.Ctx) error
	Refresh(*fiber.Ctx) error
}

type AuthHandler struct {
	service services.IAuthServices
}

func NewAuthHandler(as services.IAuthServices) AuthHandler {
	return AuthHandler{
		service: as,
	}
}

func (ah AuthHandler) Login(c *fiber.Ctx) error {
	var payload req.AuthRequest
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": errors})

	}

	response, err := ah.service.LoginUser(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah AuthHandler) Logout(c *fiber.Ctx) error {

	email := c.Locals("email").(string)

	err := ah.service.LogoutUser(email)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": "user logged out"})
}

func (ah AuthHandler) Refresh(c *fiber.Ctx) error {
	var payload req.RefreshRequest
	err := c.BodyParser(&payload)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	response, err := ah.service.RefreshToken(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

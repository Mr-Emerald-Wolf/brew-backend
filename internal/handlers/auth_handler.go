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
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	response, err := ah.service.LoginUser(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah AuthHandler) Logout() error { return nil }

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

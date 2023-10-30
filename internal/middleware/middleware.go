package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func VerifyUserToken(ctx *fiber.Ctx) error {
	var token string

	authorizationHeader := ctx.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) > 1 && fields[0] == "Bearer" {
		token = fields[1]
	}

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	res, err := services.ValidateToken(token, "SECRET")

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if res.Role != "USER" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Not a user"})
	}

	ctx.Set("currentUser", res.Id.String())
	return ctx.Next()
}

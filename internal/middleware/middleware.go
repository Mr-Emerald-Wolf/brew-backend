package middleware

import (
	"context"
	"database/sql"
	"errors"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mr-emerald-wolf/brew-backend/database"
)

// Middleware JWT function
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("ACCESS_SECRET_KEY"))},
	})
}

func CheckUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	if token == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error":  "Could Not Parse User JWT",
			"status": false,
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "malformed jwt",
			"status": false,
			"data": map[string]bool{
				"malformed": true,
			},
		})
	}
	email := claims["sub"].(string)
	user, err := database.DB.GetUserByEmail(context.Background(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user does not exist",
				"status":  false,
			})
		}
	}
	c.Locals("user", user.Uuid)
	return c.Next()
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/handlers"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func CreateAuthRoutes(app *fiber.App) {

	AuthService := services.NewAuthService(database.DB)
	AuthHandler := handlers.NewAuthHandler(AuthService)

	incomingRoutes := app.Group("/auth")
	incomingRoutes.Post("/login", AuthHandler.Login)
	incomingRoutes.Post("/refresh", AuthHandler.Refresh)

}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/internal/handlers"
	"github.com/mr-emerald-wolf/brew-backend/internal/repository"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func CreateAuthRoutes(app *fiber.App) {
	UserRepository := repository.NewUserRepository()
	AuthService := services.NewAuthService(UserRepository)
	AuthHandler := handlers.NewAuthHandler(AuthService)

	incomingRoutes := app.Group("/auth")
	incomingRoutes.Post("/login", AuthHandler.Login)
	incomingRoutes.Post("/refresh", AuthHandler.Refresh)

}

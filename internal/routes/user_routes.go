package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/internal/handlers"
	"github.com/mr-emerald-wolf/brew-backend/internal/repository"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func CreateUserRoutes(app *fiber.App) {
	UserRepository := repository.NewUserRepository()
	UserService := services.NewUserService(UserRepository)
	UserHandler := handlers.NewUserHandler(UserService)

	incomingRoutes := app.Group("/users")

	incomingRoutes.Post("/", UserHandler.NewUser)
	incomingRoutes.Get("/:uuid", UserHandler.GetUser)
	incomingRoutes.Get("/all", UserHandler.GetAllUsers)
	incomingRoutes.Patch("/:uuid", UserHandler.UpdateUser)
	incomingRoutes.Delete("/:uuid", UserHandler.DeleteUser)
}

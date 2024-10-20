package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/handlers"
	"github.com/mr-emerald-wolf/brew-backend/internal/middleware"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func CreateUserRoutes(app *fiber.App) {
	UserService := services.NewUserService(database.DB)
	UserHandler := handlers.NewUserHandler(UserService)

	incomingRoutes := app.Group("/users")

	incomingRoutes.Post("/", UserHandler.NewUser)
	incomingRoutes.Get("/all", UserHandler.GetAllUsers)
	incomingRoutes.Get("/me", middleware.VerifyUserToken, UserHandler.Me)
	incomingRoutes.Get("/:uuid", UserHandler.GetUser)
	incomingRoutes.Patch("/:uuid", UserHandler.UpdateUser)
	incomingRoutes.Delete("/:uuid", UserHandler.DeleteUser)
}

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

	incomingRoutes.Use(middleware.Protected)
	incomingRoutes.Use(middleware.CheckUser)
	
	incomingRoutes.Get("/all", UserHandler.GetAllUsers)
	incomingRoutes.Get("/me", UserHandler.Me)
	incomingRoutes.Patch("/", UserHandler.UpdateUser)
	incomingRoutes.Delete("/", UserHandler.DeleteUser)
	incomingRoutes.Get("/:uuid", UserHandler.GetUser)
}

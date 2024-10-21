package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/handlers"
	"github.com/mr-emerald-wolf/brew-backend/internal/middleware"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func CreateCoffeeRoutes(app *fiber.App) {
	CoffeeService := services.NewCoffeeService(database.DB)
	CoffeeHandler := handlers.NewCoffeeHandler(CoffeeService)
	incomingRoutes := app.Group("/coffee")

	incomingRoutes.Get("/", CoffeeHandler.GetAllCoffees)
	incomingRoutes.Get("/:uuid", CoffeeHandler.GetCoffee)

	incomingRoutes.Use(middleware.Protected())
	incomingRoutes.Use(middleware.CheckUser)

	incomingRoutes.Post("/", CoffeeHandler.NewCoffee)
	incomingRoutes.Patch("/:uuid", CoffeeHandler.UpdateCoffee)
	incomingRoutes.Delete("/:uuid", CoffeeHandler.DeleteCoffee)
}

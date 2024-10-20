package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/brew-backend/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/handlers"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
)

func CreateCoffeeRoutes(app *fiber.App) {
	CoffeeService := services.NewCoffeeService(database.DB)
	CoffeeHandler := handlers.NewCoffeeHandler(CoffeeService)

	coffeeRoutes := app.Group("/coffee")

	coffeeRoutes.Post("/", CoffeeHandler.NewCoffee)
	coffeeRoutes.Get("/", CoffeeHandler.GetAllCoffees)
	coffeeRoutes.Get("/:uuid", CoffeeHandler.GetCoffee)
	coffeeRoutes.Patch("/:uuid", CoffeeHandler.UpdateCoffee)
	coffeeRoutes.Delete("/:uuid", CoffeeHandler.DeleteCoffee)
}

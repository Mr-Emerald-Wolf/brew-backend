package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mr-emerald-wolf/brew-backend/internal/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/migrations"
	"github.com/mr-emerald-wolf/brew-backend/internal/routes"
)

func Start() {
	
	database.InitializeDB()
	migrations.RunMigrations()
	
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
	}))
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
		},
	))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "true",
			"message": "Pong!",
		})
	})

	routes.CreateUserRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})


	log.Fatal(app.Listen("localhost:8000"))
}

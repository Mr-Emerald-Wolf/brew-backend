package handlers

import (
	"github.com/gofiber/fiber/v2"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	"github.com/mr-emerald-wolf/brew-backend/internal/services"
	"github.com/mr-emerald-wolf/brew-backend/internal/utils"
)

type CoffeeHandler struct {
	service services.ICoffeeService
}

func NewCoffeeHandler(cs services.ICoffeeService) CoffeeHandler {
	return CoffeeHandler{
		service: cs,
	}
}

func (ch *CoffeeHandler) NewCoffee(c *fiber.Ctx) error {
	// Extract user ID from the request, for example, from the authentication token or session.
	// userID := "user_id_placeholder"

	var payload req.CoffeeCreateRequest
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": errors})
	}

	response, err := ch.service.CreateCoffee(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": true, "coffee": response})
	return nil
}

func (ch *CoffeeHandler) GetAllCoffees(c *fiber.Ctx) error {
	coffees, err := ch.service.FindAllCoffees()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(coffees)
}

func (ch *CoffeeHandler) GetCoffee(c *fiber.Ctx) error {
	coffeeID := c.Params("uuid") // Extract the coffee ID from the request parameters.

	coffee, err := ch.service.FindCoffee(coffeeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(coffee)
}

func (ch *CoffeeHandler) UpdateCoffee(c *fiber.Ctx) error {
	coffeeID := c.Params("uuid") // Extract the coffee ID from the request parameters.

	var payload req.CoffeeUpdateRequest
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "error": errors})
	}

	response, err := ch.service.UpdateCoffee(coffeeID, payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ch *CoffeeHandler) DeleteCoffee(c *fiber.Ctx) error {
	coffeeID := c.Params("uuid") // Extract the coffee ID from the request parameters.

	err := ch.service.DeleteCoffee(coffeeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": "Coffee Deleted"})
}

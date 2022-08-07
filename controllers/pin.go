package controllers

import (
	"travel_backend/requests"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PinController struct {
	Database *gorm.DB
}

func NewPinController(database *gorm.DB) PinController {
	return PinController{
		Database: database,
	}
}

func (p *PinController) CreatePin(c *fiber.Ctx) error {
	var data requests.PinCreate
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Pinu created successfully. 👋"})
}

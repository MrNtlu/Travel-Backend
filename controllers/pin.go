package controllers

import (
	"travel_backend/models"
	"travel_backend/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

// Create Pin
// @Summary Create Pin
// @Description Create pin
// @Tags pin
// @Accept application/json
// @Produce application/json
// @Param pincreate body requests.PinCreate true "Pin Create"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /pin/create [post]
func (p *PinController) CreatePin(c *fiber.Ctx) error {
	var data requests.PinCreate
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	pinModel := models.NewPinModel(p.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	if err := pinModel.CreatePin(data, int(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Pin created successfully. 👋"})
}

// Get Pins
// @Summary Get Pins by User ID
// @Description Get pins by user id
// @Tags pin
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /pin [get]
func (p *PinController) GetPinsByUserID(c *fiber.Ctx) error {
	pinModel := models.NewPinModel(p.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	pins, err := pinModel.GetPinsByUserID(int(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pins fetched successfully. 👋", "data": pins})
}

package controllers

import (
	"travel_backend/models"
	"travel_backend/requests"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LocationController struct {
	Database *gorm.DB
}

func NewLocationController(database *gorm.DB) LocationController {
	return LocationController{
		Database: database,
	}
}

// Get Area & City List
// @Summary Get Area and City List by Country
// @Description Returns area and city list by country
// @Tags location
// @Accept application/json
// @Produce application/json
// @Param locationcountry body requests.LocationCountry true "Location Country"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /location/area [get]
func (l *LocationController) GetAreaCityList(c *fiber.Ctx) error {
	var data requests.LocationCountry
	if err := c.QueryParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	locationModel := models.NewLocationModel(l.Database)

	areaCityList, err := locationModel.GetAreaCityList(data.Country)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Locations fetched successfully. 👋", "data": areaCityList})
}

// Get City List
// @Summary Get City List by Country
// @Description Returns city list by country
// @Tags location
// @Accept application/json
// @Produce application/json
// @Param locationcountry body requests.LocationCountry true "Location Country"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /location/city [get]
func (l *LocationController) GetCityList(c *fiber.Ctx) error {
	var data requests.LocationCountry
	if err := c.QueryParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	locationModel := models.NewLocationModel(l.Database)

	cityList, err := locationModel.GetCityList(data.Country)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "City list fetched successfully. 👋", "data": cityList})
}

// Get Country List
// @Summary Get Country List
// @Description Returns country list
// @Tags location
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /location/country [get]
func (l *LocationController) GetCountryList(c *fiber.Ctx) error {
	locationModel := models.NewLocationModel(l.Database)

	cityList, err := locationModel.GetCountryList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Country list fetched successfully. 👋", "data": cityList})
}

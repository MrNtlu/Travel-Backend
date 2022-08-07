package routes

import (
	"travel_backend/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setLocationRouter(router fiber.Router, db *gorm.DB) {
	locationController := controllers.NewLocationController(db)

	auth := router.Group("/location")

	auth.Get("/area", locationController.GetAreaCityList)
	auth.Get("/city", locationController.GetCityList)
	auth.Get("/country", locationController.GetCountryList)
}

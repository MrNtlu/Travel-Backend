package routes

import (
	"travel_backend/controllers"
	"travel_backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setPinRouter(router fiber.Router, db *gorm.DB) {
	pinController := controllers.NewPinController(db)

	pin := router.Group("/pin")

	pin.Get("/", middlewares.JWTProtected(), pinController.GetPinsByUserID)
	pin.Post("/create", middlewares.JWTProtected(), pinController.CreatePin)
}

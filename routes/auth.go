package routes

import (
	"travel_backend/controllers"
	"travel_backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setAuthRouter(router fiber.Router, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	auth := router.Group("/auth")

	auth.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Hello, World 👋!"})
	})

	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	auth.Get("/token", middlewares.JWTProtected(), controllers.TokenTest)
}

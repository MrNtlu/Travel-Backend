package routes

import (
	"travel_backend/controllers"
	"travel_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func setAuthRouter(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Hello, World 👋!"})
	})

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/token", middlewares.JWTProtected(), controllers.TokenTest)
}

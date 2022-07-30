package routes

import "github.com/gofiber/fiber/v2"

func setAuthRouter(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Hello, World 👋!"})
	})
}

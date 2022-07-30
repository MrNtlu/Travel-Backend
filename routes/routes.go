package routes

import "github.com/gofiber/fiber/v2"

func SetRoutes(router fiber.Router) {

	setAuthRouter(router)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetRoutes(router fiber.Router, db *gorm.DB) {

	setAuthRouter(router, db)
	setImageRouter(router, db)
	setLocationRouter(router, db)
}

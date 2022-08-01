package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(router fiber.Router, db *sql.DB) {

	setAuthRouter(router, db)
}

package routes

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetRoutes(router fiber.Router, db *gorm.DB, awsSession *session.Session) {
	setAuthRouter(router, db)
	setImageRouter(router, db, awsSession)
	setLocationRouter(router, db)
	setPinRouter(router, db)
}

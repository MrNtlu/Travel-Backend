package routes

import (
	"travel_backend/controllers"
	"travel_backend/middlewares"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setImageRouter(router fiber.Router, db *gorm.DB, awsSession *session.Session) {
	imageController := controllers.NewImageController(db, awsSession)

	auth := router.Group("/image")

	auth.Post("/upload", middlewares.JWTProtected(), imageController.UploadImage)
	auth.Get("/", middlewares.JWTProtected(), imageController.GetImagesByUserID)
	auth.Get("/country", middlewares.JWTProtected(), imageController.GetImagesByCountry)
	auth.Get("/location", middlewares.JWTProtected(), imageController.GetImagesByLocation)
	auth.Delete("/", middlewares.JWTProtected(), imageController.DeleteImageByID)
}

package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"travel_backend/databases"
	"travel_backend/docs"
	"travel_backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Travel Logger API
// @version 1.0
// @description REST Api of Travel Logger.

// @contact.name Burak Fidan
// @contact.email mrntlu@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	db, err := databases.SetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.Database.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	routes.SetRoutes(v1, db.Database)

	//TODO: Remove
	app.Post("/", func(c *fiber.Ctx) error {
		c.Accepts("image/png")
		c.Accepts("png")

		file, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		headerType := file.Header["Content-Type"][0]
		if headerType != "" && !strings.HasPrefix(headerType, "image") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong file type."})
		}

		image, err := file.Open()
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		defer image.Close()

		fmt.Println(headerType)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Image received 👋!"})

		// Save file to root directory:
		// return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})

	//Multiple image
	app.Post("/images", func(c *fiber.Ctx) error {
		var err error
		if form, err := c.MultipartForm(); err == nil {

			// Get all files from "documents" key:
			files := form.File["images"]
			// => []*multipart.FileHeader

			var fileMap = make(map[string]int64)
			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

				headerType := file.Header["Content-Type"][0]
				if headerType != "" && !strings.HasPrefix(headerType, "image") {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong file type."})
				}

				fileMap[file.Filename] = file.Size
			}

			fmt.Println(fileMap)

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fileMap})
		}

		return err
	})

	docs.SwaggerInfo.BasePath = "/api/v1"
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":8080")
}

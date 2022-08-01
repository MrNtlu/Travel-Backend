package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"travel_backend/databases"
	"travel_backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

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
	defer db.Database.Close()

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

	app.Listen(":8080")
}

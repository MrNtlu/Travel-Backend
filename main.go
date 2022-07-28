package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Hello, World 👋!"})
	})

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

		fmt.Println(headerType)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Image received 👋!", "size": file.Size})

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

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"travel_backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
)

func imageProcessing(buffer []byte, quality int, dirname string) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{
		Width:   1080,
		Height:  1920,
		Crop:    true,
		Quality: quality,
		Gravity: bimg.GravityCentre,
	})
	if err != nil {
		return filename, err
	}

	writeError := bimg.Write(fmt.Sprintf("./"+dirname+"/%s", filename), processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}

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

//TODO:
// Consider sending image to external source and process
// AWS S3 upload
// https://docs.gofiber.io/
// https://medium.com/wesionary-team/aws-sdk-for-go-and-uploading-a-file-using-s3-bucket-df7425317a40
// https://www.google.com/search?q=amazon+s3+image+compression&oq=amazon+s3+image+compression&aqs=edge.0.0i19j0i19i22i30l2j69i64l2.10257j0j4&sourceid=chrome&ie=UTF-8
// If authenticated get from token jwt.ExtractClaims(c)["id"].(string)
// https://github.com/gofiber/jwt
func main() {
	app := fiber.New()
	app.Use(recover.New())

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	routes.SetRoutes(v1)

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

		image, err := file.Open()
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		defer image.Close()

		buffer, err := io.ReadAll(image)
		if err != nil {
			panic(err)
		}

		errDir := createFolder("uploads")
		if errDir != nil {
			panic(errDir)
		}

		filename, err := imageProcessing(buffer, 100, "uploads")
		if err != nil {
			panic(err)
		}

		fmt.Println(headerType)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Image received 👋!", "image": "http://localhost:8080/uploads/" + filename})

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

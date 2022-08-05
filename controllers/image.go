package controllers

import (
	"fmt"
	"strings"
	"travel_backend/models"
	"travel_backend/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type ImageController struct {
	Database *gorm.DB
}

func NewImageController(database *gorm.DB) ImageController {
	return ImageController{
		Database: database,
	}
}

// Upload Image
// @Summary Upload Image
// @Description Users can upload images
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param imagecreate formData requests.ImageCreate true "Image Create"
// @Param file formData file true "Image File"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /image/upload [post]
func (i *ImageController) UploadImage(c *fiber.Ctx) error {
	var data requests.ImageCreate
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	headerType := file.Header["Content-Type"][0]
	if headerType != "" && !strings.HasPrefix(headerType, "image") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong file type."})
	}

	imageModel := models.NewImageModel(i.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	data.Image = *file
	if err := imageModel.CreateImage(data, int(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Image uploaded successfully. 👋"})
}

// Get Images
// @Summary Get Images by User ID
// @Description Returns images by user id
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param pagination body requests.Pagination true "Pagination"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /image [get]
func (i *ImageController) GetImagesByUserID(c *fiber.Ctx) error {
	var data requests.Pagination
	if err := c.QueryParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	imageModel := models.NewImageModel(i.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	images, err := imageModel.GetImagesByUserID(int(id), data.Page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Images fetched successfully. 👋", "data": images})
}

// Get Images by Country
// @Summary Get Images by Country
// @Description Returns images by user id and country
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param imagebycountry body requests.ImageByCountry true "Image by country"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /image/country [get]
func (i *ImageController) GetImagesByCountry(c *fiber.Ctx) error {
	var data requests.ImageByCountry
	if err := c.QueryParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	imageModel := models.NewImageModel(i.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	images, err := imageModel.GetImagesByCountry(int(id), data.Page, data.Country)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Images fetched successfully. 👋", "data": images})
}

// Get Images by Location
// @Summary Get Images by Location ID
// @Description Returns images by user id and location id
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param imagebylocation body requests.ImageByLocation true "Image by location id"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /image/location [get]
func (i *ImageController) GetImagesByLocation(c *fiber.Ctx) error {
	var data requests.ImageByLocation
	if err := c.QueryParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	imageModel := models.NewImageModel(i.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	images, err := imageModel.GetImagesByLocation(int(id), data.Page, data.Location)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Images fetched successfully. 👋", "data": images})
}

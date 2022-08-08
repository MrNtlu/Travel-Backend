package controllers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"travel_backend/models"
	"travel_backend/requests"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageController struct {
	Database   *gorm.DB
	AWSSession *session.Session
}

func NewImageController(database *gorm.DB, session *session.Session) ImageController {
	return ImageController{
		Database:   database,
		AWSSession: session,
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

	fileHeader, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	headerType := fileHeader.Header["Content-Type"][0]
	if headerType != "" && !strings.HasPrefix(headerType, "image") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong file type."})
	}

	imageModel := models.NewImageModel(i.Database)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	data.Image = *fileHeader

	uploader := s3manager.NewUploader(i.AWSSession)
	bucket := os.Getenv("AWS_BUCKET_NAME")

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	awsKey := strconv.Itoa(int(id)) + "/" + uuid.NewString() + data.Image.Filename

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(awsKey),
		Body:   file,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image. " + err.Error()})
	}

	if err := imageModel.CreateImage(data, awsKey, upload.Location, int(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Image uploaded successfully. 👋", "upload_location": aws.StringValue(&upload.Location)})
}

// Get Images
// @Summary Get Images by User ID
// @Description Returns images by user id
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param pagination query requests.Pagination true "Pagination"
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
// @Param imagebycountry query requests.ImageByCountry true "Image by country"
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
// @Param imagebylocation query requests.ImageByLocation true "Image by location id"
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

// Delete Image by ID
// @Summary Delete Image by ID
// @Description Returns image by user id and image id
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param imagebylocation body requests.ImageByLocation true "Image by location id"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /image [delete]
func (i *ImageController) DeleteImageByID(c *fiber.Ctx) error {
	var data requests.ID
	if err := c.BodyParser(&data); err != nil {
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

	image, err := imageModel.GetImageByID(int(id), data.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if err := imageModel.DeleteImageByID(int(id), data.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	bucket := os.Getenv("AWS_BUCKET_NAME")
	awsManager := s3.New(i.AWSSession)

	if _, err = awsManager.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(image.AWSImageKey),
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Image deleted successfully."})
}

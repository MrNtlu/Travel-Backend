package controllers

import (
	"os"
	"strconv"
	"time"
	"travel_backend/models"
	"travel_backend/requests"
	"travel_backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	Database *gorm.DB
}

func NewAuthController(database *gorm.DB) AuthController {
	return AuthController{
		Database: database,
	}
}

// Login
// @Summary User Login
// @Description Allows users to login and generate jwt token
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param register body requests.Login true "User login credentials"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /auth/login [post]
func (a *AuthController) Login(c *fiber.Ctx) error {
	var data requests.Login
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	userModel := models.NewUserModel(a.Database)
	user, err := userModel.FindUserByEmail(data.EmailAddress)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if err := utils.CheckPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect email or password."})
	}

	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"id":       user.ID,
		"username": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Welcome back " + user.FirstName + " 👋", "token": jwtToken})
}

// Register
// @Summary User Registration
// @Description Allows users to register
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param register body requests.Register true "User registration info"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Router /auth/register [post]
func (a *AuthController) Register(c *fiber.Ctx) error {
	var data requests.Register
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	userModel := models.NewUserModel(a.Database)
	if err := userModel.CreateUser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Registered 👋", "email": data.EmailAddress})
}

func TokenTest(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	username := claims["username"].(string)

	return c.SendString("Welcome " + strconv.FormatUint(uint64(id), 10) + " " + username)
}

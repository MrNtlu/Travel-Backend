package controllers

import (
	"fmt"
	"os"
	"time"
	"travel_backend/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct {
}

func Login(c *fiber.Ctx) error {
	var loginBody requests.Login
	if err := c.BodyParser(&loginBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(loginBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Throws Unauthorized error
	if loginBody.EmailAddress != "burak@gmail.com" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"id":   "burak@gmail.com",
		"name": "Burak Fidan",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Logged in 👋", "token": jwtToken})
}

func Register(c *fiber.Ctx) error {
	var registerBody requests.Register
	if err := c.BodyParser(&registerBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := validateStruct(registerBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Registered 👋", "email": registerBody.EmailAddress})
}

func TokenTest(c *fiber.Ctx) error {
	fmt.Println(c.Locals("user"))
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return c.SendString("Welcome " + id)
}

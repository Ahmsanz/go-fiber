package login

import (
	"go-fiber/domain/models"
	"go-fiber/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginHandler struct{}

func RegisterLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (lh *LoginHandler) Login(c *fiber.Ctx) error {
	var credentials models.Credentials

	if err := c.BodyParser(&credentials); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	username := credentials.Username
	password := credentials.Password

	// Throws Unauthorized error
	if username == "" || password == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "the token is absent or expired",
		})
	}

	expiration := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := jwt.MapClaims{
		"username": username,
		"isValid":  true,
		"exp":      expiration,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(utils.Config("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"error":   nil,
		"session": models.BuildSession(username, t, time.Duration(expiration)),
	})
}

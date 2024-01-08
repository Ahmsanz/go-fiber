package routes

import (
	"go-fiber/application/use-cases/login"
	"go-fiber/utils"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Post("/login", login.RegisterLoginHandler().Login)

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(utils.Config("JWT_SECRET"))},
	}))

	// books routes
	RegisterBooksRoutes(api)
}

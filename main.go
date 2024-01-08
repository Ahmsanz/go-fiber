package main

import (
	"fmt"
	"go-fiber/application/routes"
	"go-fiber/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := utils.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You're reaching the go-fiber-server",
			"data":    nil,
		})
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", utils.Config("PORT"))))
}

package main

import (
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
		return c.SendString("Hello, World!")
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":6000"))
}

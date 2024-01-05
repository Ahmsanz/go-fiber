package routes

import (
	"go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBooksRoutes(api fiber.Router) fiber.Router {
	books := api.Group("/books")

	books.Post("/", handlers.AddBook)
	books.Get("/", handlers.GetBooks)
	books.Get("/:bookID", handlers.GetBook)
	books.Patch("/:bookID", handlers.UpdateBook)
	books.Delete("/:bookID", handlers.DeleteBook)

	return api
}

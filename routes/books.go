package routes

import (
	addBook "go-fiber/application/use-cases/add-book"
	listBooks "go-fiber/application/use-cases/list-books"
	"go-fiber/infrastructure/repositories"

	"github.com/gofiber/fiber/v2"
)

func RegisterBooksRoutes(api fiber.Router) fiber.Router {
	books := api.Group("/books")

	books.Post("/", addBook.RegisterAddBookHandler(repositories.BuildPostgresRepo()).AddBook)
	books.Get("/", listBooks.RegisterListBookHandler(repositories.BuildPostgresRepo()).ListBooks)
	// books.Get("/:bookID", handlers.GetBook)
	// books.Patch("/:bookID", handlers.UpdateBook)
	// books.Delete("/:bookID", handlers.DeleteBook)

	return api
}

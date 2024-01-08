package routes

import (
	addBook "go-fiber/application/use-cases/add-book"
	deleteBook "go-fiber/application/use-cases/delete-book"
	getBook "go-fiber/application/use-cases/get-book"
	listBooks "go-fiber/application/use-cases/list-books"
	updateBook "go-fiber/application/use-cases/update-book"
	"go-fiber/infrastructure/repositories"

	"github.com/gofiber/fiber/v2"
)

func RegisterBooksRoutes(api fiber.Router) fiber.Router {
	books := api.Group("/books")

	books.Post("/", addBook.RegisterAddBookHandler(repositories.BuildPostgresRepo()).AddBook)
	books.Get("/", listBooks.RegisterListBookHandler(repositories.BuildPostgresRepo()).ListBooks)
	books.Get("/:bookID", getBook.RegisterGetBookHandler(repositories.BuildPostgresRepo()).GetBook)
	books.Patch("/:bookID", updateBook.RegisterUpdateBookHandler(repositories.BuildPostgresRepo()).UpdateBook)
	books.Delete("/:bookID", deleteBook.RegisterDeleteBookHandler(repositories.BuildPostgresRepo()).SoftDeleteBook)

	return api
}

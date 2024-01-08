package addBook

import (
	"go-fiber/domain/interfaces"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type AddBookHandler struct {
	repo interfaces.BookRepo
}

func RegisterAddBookHandler(r interfaces.BookRepo) *AddBookHandler {
	return &AddBookHandler{repo: r}
}

func (ah AddBookHandler) AddBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something failed created the new book",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	result, err := ah.repo.Create(book)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something failed created the new book",
			"data":    nil,
			"error":   result,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "Book successfully created",
		"data":    &book,
	})
}

// func GetBook(c *fiber.Ctx) error {
// 	return Repo.ReadOne(c)
// }

// func UpdateBook(c *fiber.Ctx) error {
// 	return Repo.Update(c)
// }

// func DeleteBook(c *fiber.Ctx) error {
// 	return Repo.Delete(c)
// }

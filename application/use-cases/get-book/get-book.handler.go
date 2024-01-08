package getBook

import (
	"go-fiber/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type GetBookHandler struct {
	repo interfaces.BookRepo
}

func RegisterGetBookHandler(r interfaces.BookRepo) *GetBookHandler {
	return &GetBookHandler{repo: r}
}

func (gh *GetBookHandler) GetBook(c *fiber.Ctx) error {
	book, err := gh.repo.ReadOne(c.Params("bookID"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something failed getting the book",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Book found",
		"data":    &book,
	})
}

package listBooks

import (
	"go-fiber/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type ListBooksHandler struct {
	repo interfaces.BookRepo
}

func RegisterListBookHandler(r interfaces.BookRepo) *ListBooksHandler {
	return &ListBooksHandler{repo: r}
}

func (lh ListBooksHandler) ListBooks(c *fiber.Ctx) error {
	result, err := lh.repo.Read()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong listing the books",
			"data":    nil,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "Books found",
		"data":    &result,
	})
}

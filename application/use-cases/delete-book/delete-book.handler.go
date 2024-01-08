package deleteBook

import (
	"go-fiber/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type DeleteBookHandler struct {
	repo interfaces.BookRepo
}

func RegisterDeleteBookHandler(r interfaces.BookRepo) *DeleteBookHandler {
	return &DeleteBookHandler{repo: r}
}

func (dh *DeleteBookHandler) SoftDeleteBook(c *fiber.Ctx) error {
	if err := dh.repo.Delete(c.Params("bookID")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "error deleting book",
			"data":    nil,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"success": false,
		"message": "book sucessfully deleted",
		"data":    nil,
		"error":   nil,
	})
}

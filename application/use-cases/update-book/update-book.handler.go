package updateBook

import (
	"go-fiber/domain/interfaces"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateBookHandler struct {
	repo interfaces.BookRepo
}

func RegisterUpdateBookHandler(r interfaces.BookRepo) *UpdateBookHandler {
	return &UpdateBookHandler{repo: r}
}

func (uh *UpdateBookHandler) UpdateBook(c *fiber.Ctx) error {
	payload := new(models.Book)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "error parsing body",
			"data":    nil,
			"error":   err,
		})
	}

	result, err := uh.repo.Update(c.Params("bookID"), payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "error updating book",
			"data":    nil,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": false,
		"message": "book updated successfully",
		"data":    result,
		"error":   nil,
	})
}

package handlers

import (
	"go-fiber/models"
	"go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

var Books []models.Book

func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something failed created the new book",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	result := utils.DB.Create(&book)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something failed created the new book",
			"data":    nil,
			"error":   result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "Book successfully created",
		"data":    book,
	})
}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if err := utils.DB.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong",
			"data":    nil,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Found books",
		"data":    &books,
	})
}

func GetBook(c *fiber.Ctx) error {
	var b models.Book

	if err := utils.DB.First(&b, "id = ?", c.Params("bookID")).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong",
			"data":    nil,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Found book",
		"data":    &b,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	var b models.Book

	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	if err := utils.DB.First(&b, "id = ?", c.Params("bookID")).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	if err := utils.DB.Model(&b).Updates(c.BodyParser(&b)).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Book successfully updated",
		"data":    &b,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	var book models.Book

	bookId := c.Params("bookID")

	if err := utils.DB.Delete(&book, "id = ?", bookId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"success": true,
		"message": "Book successfully deleted",
		"data":    nil,
	})
}

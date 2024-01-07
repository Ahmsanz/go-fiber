package interfaces

import "go-fiber/domain/models"

type BookRepo interface {
	Create(payload *models.Book) (*models.Book, error)
	Read() (*[]models.Book, error)
	ReadOne(id string) (*models.Book, error)
	Update(id string, payload models.Book) (*models.Book, error)
	Delete(id string) error
}

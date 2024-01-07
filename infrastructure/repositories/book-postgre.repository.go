package repositories

import (
	"go-fiber/domain/models"
	"go-fiber/utils"

	"gorm.io/gorm"
)

type PostgresBookRepo struct {
	db *gorm.DB
}

func BuildPostgresRepo() PostgresBookRepo {
	return PostgresBookRepo{db: utils.PostgresDB}
}

func (repo PostgresBookRepo) Create(book *models.Book) (*models.Book, error) {
	if err := repo.db.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (repo PostgresBookRepo) Read() (*[]models.Book, error) {
	var books []models.Book

	if err := repo.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return &books, nil
}

func (repo PostgresBookRepo) ReadOne(id string) (*models.Book, error) {
	var b models.Book

	if err := repo.db.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (repo PostgresBookRepo) Update(id string, book models.Book) (*models.Book, error) {

	if err := repo.db.First(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := repo.db.Model(&book).Updates(book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo PostgresBookRepo) Delete(id string) error {
	if err := repo.db.Delete(&models.Book{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

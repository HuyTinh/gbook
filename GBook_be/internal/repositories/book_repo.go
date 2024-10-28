package repositories

import (
	"GBook_be/internal/models"

	"gorm.io/gorm"
)

type (
	BookRepository interface {
		FindAllBook() ([]models.Book, error)

		// FindBookById(bookId int64)

		// FindBookBySlug(bookSlug string)

		// SaveBook(newBook models.Book)

		// UpdateBook(bookId int64, updateBook models.Book)

		// DeleteBook(bookId int64)
	}

	BookRepositoryImpl struct {
		db *gorm.DB
	}
)

func (br BookRepositoryImpl) FindAllBook() ([]models.Book, error) {
	var books []models.Book
	result := br.db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil

}

func ProvideBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

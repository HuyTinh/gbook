package authors

import (
	"GBook_be/internal/models"

	"gorm.io/gorm"
)

type (
	AuthorRepository interface {
		FindAllAuthor() ([]models.Author, error)

		// FindBookById(bookId int64)

		// FindBookBySlug(bookSlug string)

		// SaveBook(newBook models.Book)

		// UpdateBook(bookId int64, updateBook models.Book)

		// DeleteBook(bookId int64)
	}

	AuthorRepositoryImpl struct {
		db *gorm.DB
	}
)

func (ar AuthorRepositoryImpl) FindAllAuthor() ([]models.Author, error) {
	var authors []models.Author
	result := ar.db.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil

}

func ProvideAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{
		db: db,
	}
}

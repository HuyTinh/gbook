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

		SaveAuthor(newAuthor models.Author)

		// UpdateBook(bookId int64, updateBook models.Book)

		// DeleteBook(bookId int64)
	}

	AuthorRepositoryImpl struct {
		db *gorm.DB
	}
)

func (ar AuthorRepositoryImpl) FindAllAuthor() ([]models.Author, error) {
	var authors []models.Author
	result := ar.db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil

}

func (ar AuthorRepositoryImpl) SaveAuthor(saveAuthor models.Author) {

	var author models.Author
	if err := ar.db.Find(&author, saveAuthor.ID).Error; err != nil {
		return
	}

	if err := ar.db.Create(&saveAuthor).Error; err != nil {
		return
	}
}

func ProvideAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{
		db: db,
	}
}

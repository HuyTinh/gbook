package books

import (
	"GBook_be/internal/models"

	"gorm.io/gorm"
)

type (
	BookRepository interface {
		FindAllBook() ([]models.Book, error)

		// FindBookById(bookId int64)

		// FindBookBySlug(bookSlug string)

		SaveBook(saveBook models.Book)

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

func (br BookRepositoryImpl) SaveBook(saveBook models.Book) {

	var book models.Book
	if err := br.db.Find(&book, saveBook.ID).Error; err != nil {
		return
	}

	if err := br.db.Create(&saveBook).Error; err != nil {
		return
	}
}

func ProvideBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

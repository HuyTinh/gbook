package books

import (
	"GBook_be/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	BookRepository interface {
		FindAllBook() ([]models.Book, error)

		FindBookById(bookId int64) (models.Book, error)

		FindBookBySlug(bookSlug string) (models.Book, error)

		SaveBook(saveBook models.Book) (models.Book, error)

		// UpdateBook(bookId int64, updateBook models.Book)

		// DeleteBook(bookId int64)
	}

	BookRepositoryImpl struct {
		db *gorm.DB
	}
)

func ProvideBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (br BookRepositoryImpl) FindAllBook() ([]models.Book, error) {
	var books []models.Book
	result := br.db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil

}

func (br BookRepositoryImpl) SaveBook(saveBook models.Book) (models.Book, error) {

	var book models.Book

	if err := br.db.First(&book, saveBook.ID).Error; err != nil {
		return models.Book{}, err
	}

	if err := br.db.Create(&saveBook).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (br BookRepositoryImpl) FindBookById(bookId int64) (models.Book, error) {
	var book models.Book

	if err := br.db.Preload("Author").First(&book, bookId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Book{}, fmt.Errorf("book with ID %d not found", bookId)
		}

		return models.Book{}, fmt.Errorf("failed to retrieve book: %v", err)
	}

	return book, nil
}

func (br BookRepositoryImpl) FindBookBySlug(bookSlug string) (models.Book, error) {
	var book models.Book

	if err := br.db.Preload("Books").Where("slug = ?", bookSlug).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Book{}, fmt.Errorf("book with Slug %s not found", bookSlug)
		}

		return models.Book{}, fmt.Errorf("failed to retrieve book: %v", err)
	}

	return book, nil
}

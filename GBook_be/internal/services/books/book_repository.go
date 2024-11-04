package books

import (
	"GBook_be/internal/models"
	"errors"
	"fmt"
	"sync"

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
	batchSize := 512 * 3
	offset := 0

	var books []models.Book
	var wg sync.WaitGroup

	results := make(chan []models.Book)
	done := make(chan struct{})

	go func() {
		for {
			batch := make([]models.Book, 0, batchSize)
			if err := br.db.Preload("Author").Limit(batchSize).Offset(offset).Find(&batch).Error; err != nil {
				close(done)
				return
			}

			if len(batch) == 0 {
				break
			}

			results <- batch
			offset += batchSize
		}
		close(results)
	}()

	go func() {
		for batch := range results {
			books = append(books, batch...)
		}
		close(done)
	}()

	wg.Wait()
	<-done

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

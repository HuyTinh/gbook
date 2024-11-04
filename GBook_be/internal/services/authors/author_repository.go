package authors

import (
	"GBook_be/internal/models"
	"sync"

	"gorm.io/gorm"
)

type (
	AuthorRepository interface {
		FindAllAuthor() ([]models.Author, error)

		FindAuthorById(authorId int64) (models.Author, error)

		FindAuthorByName(authorName string) (models.Author, error)

		SaveAuthor(newAuthor models.Author) (models.Author, error)

		// UpdateAuthor(authorId int64, updateAuthor models.Author)

		// DeleteAuthor(authorId int64)
	}

	AuthorRepositoryImpl struct {
		db *gorm.DB
	}
)

func ProvideAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{
		db: db,
	}
}

func (ar AuthorRepositoryImpl) FindAllAuthor() ([]models.Author, error) {
	batchSize := 512 * 3

	var authors []models.Author
	offset := 0
	var wg sync.WaitGroup
	results := make(chan []models.Author)
	done := make(chan struct{})

	go func() {
		for {
			batch := make([]models.Author, 0, batchSize)
			if err := ar.db.Preload("Books").Limit(batchSize).Offset(offset).Find(&batch).Error; err != nil {
				close(done)
				return
			}

			if len(batch) == 0 {
				break // Exit if no more authors
			}

			results <- batch
			offset += batchSize
		}
		close(results)
	}()

	go func() {
		for batch := range results {
			authors = append(authors, batch...)
		}
		close(done)
	}()

	// Wait for all batches to be processed
	wg.Wait()
	<-done // Wait for the results to finish

	return authors, nil
}

func (ar AuthorRepositoryImpl) FindAuthorById(authorId int64) (models.Author, error) {
	var author models.Author

	if err := ar.db.Preload("Books").Find(&author, authorId).Error; err != nil {
		return models.Author{}, err
	}

	return author, nil
}

func (ar AuthorRepositoryImpl) FindAuthorByName(authorName string) (models.Author, error) {
	var author models.Author

	if err := ar.db.Preload("Books").Where("name = ?", authorName).First(&author).Error; err != nil {
		return models.Author{}, err
	}

	return author, nil
}

func (ar AuthorRepositoryImpl) SaveAuthor(saveAuthor models.Author) (models.Author, error) {

	var author models.Author

	if err := ar.db.Find(&author, saveAuthor.ID).Error; err != nil {
		return models.Author{}, err
	}

	if err := ar.db.Create(&saveAuthor).Error; err != nil {
		return models.Author{}, err
	}

	return author, nil
}

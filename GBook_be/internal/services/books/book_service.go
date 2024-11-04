package books

import (
	APIResponse "GBook_be/internal/dto/response"
	"GBook_be/internal/models"

	"github.com/gin-gonic/gin"
)

type BookService struct {
	repository BookRepository
}

func ProvideBookService(bookRepository BookRepository) BookService {
	return BookService{
		repository: bookRepository,
	}
}

func (bs *BookService) GetAllBook(c *gin.Context) {

	result, err := bs.repository.FindAllBook()

	if err != nil {
		c.JSON(500, APIResponse.InitializeAPIResponse(500, "Error fetching books: "+err.Error(), ""))
		return
	}

	if len(result) == 0 {
		c.JSON(200, APIResponse.InitializeAPIResponse(200, "No books found", ""))
		return
	}

	bookResponsesChan := make(chan APIResponse.BookResponse)
	done := make(chan bool)

	go func() {
		for _, book := range result {
			bookResponse := APIResponse.BookResponse{
				ID:            book.ID,
				Name:          book.Name,
				Slug:          book.Slug,
				Price:         book.Price,
				StockQuantity: book.StockQuantity,
				PublishedDate: book.PublishedDate,
				ISBN:          book.ISBN,
				Description:   book.Description,
				CoverImageURL: book.CoverImageURL,
				Author:        book.Author,
				Genre:         book.Genre,
			}

			bookResponsesChan <- bookResponse
		}
		close(bookResponsesChan)
	}()

	bookResponses := make([]APIResponse.BookResponse, 0, len(result))

	go func() {
		for bookResponse := range bookResponsesChan {
			bookResponses = append(bookResponses, bookResponse)
		}
		done <- true
	}()

	<-done

	c.JSON(200, APIResponse.InitializeAPIResponse(200, "", bookResponses))
}

func (bs *BookService) SaveBook(c *gin.Context) {

	var saveBook models.Book

	if err := c.ShouldBindJSON(&saveBook); err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, err.Error(), ""))
		return
	}

	successSaveBook, err := bs.repository.SaveBook(saveBook)

	if err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, err.Error(), ""))
		return
	}

	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", successSaveBook))
}

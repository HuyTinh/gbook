package books

import (
	APIResponse "GBook_be/internal/dto/response"
	"GBook_be/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/thoas/go-funk"
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
		c.JSON(400, APIResponse.InitializeAPIResponse(400, err.Error(), ""))
		return
	}

	bookResponses := funk.Map(result, func(book models.Book) APIResponse.BookResponse {
		var bookResponse APIResponse.BookResponse
		copier.Copy(&bookResponse, &book)
		return bookResponse
	}).([]APIResponse.BookResponse)

	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", bookResponses))
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

package authors

import (
	response "GBook_be/internal/dto/response"
	"GBook_be/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/thoas/go-funk"
)

type AuthorService struct {
	repository AuthorRepository
}

func ProvideAuthorService(bookRepository AuthorRepository) AuthorService {
	return AuthorService{
		repository: bookRepository,
	}
}

func (bc *AuthorService) GetAllAuthor(c *gin.Context) {

	result, err := bc.repository.FindAllAuthor()

	if err != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	authorResponses := funk.Map(result, func(author models.Author) response.AuthorResponse {
		var authorResponse response.AuthorResponse
		copier.Copy(&authorResponse, &author)
		return authorResponse
	}).([]response.AuthorResponse)

	c.JSON(200, response.InitializeAPIResponse(1000, "", authorResponses))
}

func (as *AuthorService) SaveAuthor(c *gin.Context) {

	var saveAuthor models.Author

	if err := c.ShouldBindJSON(&saveAuthor); err != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}
	as.repository.SaveAuthor(saveAuthor)
	c.JSON(200, response.InitializeAPIResponse(1000, "", ""))
}

// func (bc *BookService) GetBookById(c *gin.Context) {

// 	var book models.Book

// 	if result := bc.db.Preload("Author").First(&book, c.Param("id")); result.Error != nil {
// 		c.JSON(400, response.InitializeAPIResponse(400, fmt.Sprintf("Book with id = %s is not found", c.Param("id")), ""))
// 		return
// 	}

// 	var bookResponse response.BookResponse

// 	_ = copier.Copy(&bookResponse, &book)

// 	c.JSON(200, response.InitializeAPIResponse(1000, "", bookResponse))
// }

// func (bc *BookService) CreateBook(c *gin.Context) {

// 	var newBook models.Book

// 	if err := c.ShouldBindJSON(&newBook); err != nil {
// 		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
// 		return
// 	}

// 	var author models.Author
// 	if err := bc.db.Find(&author, newBook.AuthorID).Error; err != nil {
// 		c.JSON(500, response.InitializeAPIResponse(500, "Author not found", ""))
// 		return
// 	}

// 	if err := bc.db.Create(&newBook).Error; err != nil {
// 		c.JSON(500, response.InitializeAPIResponse(500, "Failed to create book", ""))
// 		return
// 	}

// 	c.JSON(200, response.InitializeAPIResponse(1000, "", newBook))
// }

// func (bc *BookService) UpdateBook(c *gin.Context) {

// 	var updateBook models.Book

// 	if err := c.ShouldBindJSON(&updateBook); err != nil {
// 		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
// 		return
// 	}

// 	if err := bc.db.Updates(&updateBook); err != nil {
// 		c.JSON(500, response.InitializeAPIResponse(500, "Failed to update book", ""))
// 		return
// 	}

// 	c.JSON(200, response.InitializeAPIResponse(200, "", updateBook))

// }

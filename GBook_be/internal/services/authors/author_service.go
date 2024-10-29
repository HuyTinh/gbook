package authors

import (
	APIResponse "GBook_be/internal/dto/response"
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
		c.JSON(400, APIResponse.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	authorResponses := funk.Map(result, func(author models.Author) APIResponse.AuthorResponse {
		var authorResponse APIResponse.AuthorResponse
		copier.Copy(&authorResponse, &author)
		return authorResponse
	}).([]APIResponse.AuthorResponse)

	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", authorResponses))
}

func (as *AuthorService) SaveAuthor(c *gin.Context) {

	var saveAuthor models.Author

	if err := c.ShouldBindJSON(&saveAuthor); err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}
	as.repository.SaveAuthor(saveAuthor)
	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", ""))
}

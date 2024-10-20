package controllers

import (
	"GBook_be/internal/dto/response"
	"GBook_be/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type BookController struct {
	db *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		db: db,
	}
}

func (bc *BookController) GetAllBooks(c *gin.Context) {

	var books []models.Book

	if result := bc.db.Preload("Author").Find(&books); result.Error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	bookResponses := funk.Map(books, func(book models.Book) response.BookResponse {
		var bookResponse response.BookResponse
		copier.Copy(&bookResponse, &book)
		return bookResponse
	}).([]response.BookResponse)

	c.JSON(200, response.InitializeAPIResponse(1000, "", bookResponses))
}

func (bc *BookController) GetBookById(c *gin.Context) {

	var book models.Book

	if result := bc.db.Preload("Author").First(&book, c.Param("id")); result.Error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, fmt.Sprintf("Book with id = %s is not found", c.Param("id")), ""))
		return
	}

	var bookResponse response.BookResponse

	_ = copier.Copy(&bookResponse, &book)

	c.JSON(200, response.InitializeAPIResponse(1000, "", bookResponse))
}

func (bc *BookController) CreateBook(c *gin.Context) {

	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	var author models.Author
	if err := bc.db.Find(&author, newBook.AuthorID).Error; err != nil {
		c.JSON(500, response.InitializeAPIResponse(500, "Author not found", ""))
		return
	}

	if err := bc.db.Create(&newBook).Error; err != nil {
		c.JSON(500, response.InitializeAPIResponse(500, "Failed to create book", ""))
		return
	}

	c.JSON(200, response.InitializeAPIResponse(1000, "", newBook))
}

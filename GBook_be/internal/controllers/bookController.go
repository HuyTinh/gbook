package controllers

import (
	"GBook_be/internal/dto/response"
	"GBook_be/internal/models"

	"github.com/gin-gonic/gin"
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

	if result := bc.db.Find(&books); result.Error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	c.JSON(200, response.InitializeAPIResponse(1000, "", books))
}

func (bc *BookController) GetBookById(c *gin.Context) {

	var books []models.Book

	if result := bc.db.First(&books, c.Param("id")); result.Error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	c.JSON(200, response.InitializeAPIResponse(1000, "", books))
}

func (bc *BookController) CreateBook(c *gin.Context) {

	var newBook models.Book

	if error := c.ShouldBindJSON(&newBook); error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	if result := bc.db.Create(&newBook); result.Error != nil {
		c.JSON(500, response.InitializeAPIResponse(500, "Failed to create book", ""))
		return
	}

	c.JSON(200, response.InitializeAPIResponse(1000, "", newBook))
}

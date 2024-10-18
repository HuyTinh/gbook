package controllers

import (
	"GBook_be/internal/database"
	"GBook_be/internal/dto/response"
	"GBook_be/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {

	var books []models.Book

	if result := database.DB.Find(&books); result.Error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	c.JSON(200, response.InitializeAPIResponse(1000, "", books))
}

func CreateBook(c *gin.Context) {

	var newBook models.Book

	if error := c.ShouldBindJSON(&newBook); error != nil {
		c.JSON(400, response.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}

	if result := database.DB.Create(&newBook); result.Error != nil {
		c.JSON(500, response.InitializeAPIResponse(500, "Failed to create book", ""))
		return
	}

	c.JSON(200, response.InitializeAPIResponse(1000, "", newBook))
}

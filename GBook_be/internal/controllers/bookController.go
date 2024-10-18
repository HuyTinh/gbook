package controllers

import (
	"GBook_be/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	books []models.Book
}

// NewBookHandler creates a new instance of BookHandler
func NewBookHandler() *BookHandler {
	return &BookHandler{
		books: []models.Book{
			{ID: 1, Title: "1984", Isbn: "1234567890", Price: 10000},
			{ID: 2, Title: "To Kill a Mockingbird", Isbn: "1234567891", Price: 20000},
		},
	}
}

func (bh *BookHandler) GetAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, bh.books)
}

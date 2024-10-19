package routes

import (
	"GBook_be/internal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBookRouter(r *gin.Engine, db *gorm.DB) {

	bookController := controllers.NewBookController(db)

	r.GET("/books", bookController.GetAllBooks)
	r.GET("/books/:id", bookController.GetBookById)
	r.POST("/books", bookController.CreateBook)
}

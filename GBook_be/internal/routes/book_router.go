package routes

import (
	"GBook_be/internal/controllers"
	"GBook_be/internal/repositories"
	"GBook_be/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBookRouter(r *gin.Engine, db *gorm.DB) {

	bookService := services.InitializeBookService(repositories.ProvideBookRepository(db))

	bookGroupRouter := r.Group("/books")

	controllers.RegisterController(bookGroupRouter, controllers.BookController(bookService))
}

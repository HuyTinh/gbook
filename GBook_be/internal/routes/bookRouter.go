package routes

import (
	"GBook_be/internal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBookRouter(r *gin.Engine, db *gorm.DB) {

	bookController := controllers.NewBookController(db)

	bookGroupRouter := r.Group("/books")
	bookGroupRouter.GET("", bookController.GetAllBooks)
	bookGroupRouter.GET("/:id", bookController.GetBookById)
	bookGroupRouter.POST("", bookController.CreateBook)
}

package routes

import (
	"GBook_be/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRouter(r *gin.Engine) {
	r.GET("/books", controllers.GetAllBooks)
	r.POST("/books", controllers.CreateBook)
}

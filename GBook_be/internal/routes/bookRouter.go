package routes

import (
	"GBook_be/internal/controllers"
	HttpMethod "GBook_be/internal/enums"
	"GBook_be/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EndPoints(controller *controllers.BookController) []utils.EndPoint {
	return []utils.EndPoint{
		{
			Method:     HttpMethod.GET,
			Path:       "",
			Controller: controller.GetAllBooks,
		},
		{
			Method:     HttpMethod.POST,
			Path:       "",
			Controller: controller.CreateBook,
		},
		{
			Method:     HttpMethod.GET,
			Path:       "/:id",
			Controller: controller.GetBookById,
		},
	}
}

func RegisterBookRouter(r *gin.Engine, db *gorm.DB) {

	bookController := controllers.InitializeBookController(db)

	bookGroupRouter := r.Group("/books")

	utils.RegisterEndPoint(bookGroupRouter, EndPoints(bookController))
}

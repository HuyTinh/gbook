package books

import (
	HTTPMethod "GBook_be/internal/enums"
	"GBook_be/internal/server"

	"github.com/gin-gonic/gin"
)

func ProvideBookController(service BookService, routerGroup *gin.RouterGroup) server.Controller {
	return server.Controller{
		RouterGroup: routerGroup,
		Routes: []server.Route{
			{
				Method:     HTTPMethod.GET,
				Path:       "",
				Controller: service.GetAllBooks,
			},
		},
	}
}

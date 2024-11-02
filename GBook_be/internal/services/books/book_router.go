package books

import (
	"os"

	"github.com/gin-gonic/gin"
)

type BookRoute struct {
	route *gin.RouterGroup
}

func ProvideBookRouter(routerGroup *gin.Engine) BookRoute {

	return BookRoute{
		route: routerGroup.Group(os.Getenv("BOOK_END_POINT")),
	}

}

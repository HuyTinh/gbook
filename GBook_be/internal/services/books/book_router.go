package books

import (
	"github.com/gin-gonic/gin"
)

type BookRoute struct {
	route *gin.RouterGroup
}

func ProvideBookRouter(routerGroup *gin.Engine) BookRoute {

	return BookRoute{
		route: routerGroup.Group("/books"),
	}
}

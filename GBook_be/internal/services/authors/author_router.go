package authors

import (
	"github.com/gin-gonic/gin"
)

type AuthorRoute struct {
	route *gin.RouterGroup
}

func ProvideAuthorRouter(routerGroup *gin.Engine) AuthorRoute {
	return AuthorRoute{
		route: routerGroup.Group("/authors"),
	}
}

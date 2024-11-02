package authors

import (
	"os"

	"github.com/gin-gonic/gin"
)

type AuthorRoute struct {
	route *gin.RouterGroup
}

func ProvideAuthorRouter(routerGroup *gin.Engine) AuthorRoute {

	return AuthorRoute{
		route: routerGroup.Group(os.Getenv("AUTHOR_END_POINT")),
	}

}

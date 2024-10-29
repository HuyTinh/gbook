package authors

import (
	"github.com/gin-gonic/gin"
)

func ProvideAuthorRouter(routerGroup *gin.Engine) *gin.RouterGroup {

	return routerGroup.Group("/authors")
}

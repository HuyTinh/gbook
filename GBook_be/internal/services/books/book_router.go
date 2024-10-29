package books

import (
	"github.com/gin-gonic/gin"
)

func ProvideBookRouter(routerGroup *gin.Engine) *gin.RouterGroup {

	return routerGroup.Group("/books")
}

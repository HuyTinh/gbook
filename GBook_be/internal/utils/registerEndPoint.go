package utils

import (
	HttpMethod "GBook_be/internal/enums"

	"github.com/gin-gonic/gin"
)

type EndPoint struct {
	Method     string
	Path       string
	Controller func(c *gin.Context)
}

func RegisterEndPoint(gRouter *gin.RouterGroup, endPoints []EndPoint) {
	for _, value := range endPoints {
		switch value.Method {
		case HttpMethod.GET:
			gRouter.GET(value.Path, value.Controller)
		case HttpMethod.POST:
			gRouter.POST(value.Path, value.Controller)
		case HttpMethod.PUT:
			gRouter.PUT(value.Path, value.Controller)
		case HttpMethod.DELETE:
			gRouter.DELETE(value.Path, value.Controller)
		}
	}
}

package server

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method     string
	Path       string
	Controller func(c *gin.Context)
}

type Controller struct {
	RouterGroup *gin.RouterGroup
	Routes      []Route
}

func ProvideController(controllers []Controller) {
	// Ánh xạ phương thức HTTP tới các hàm xử lý tương ứng

	for _, controller := range controllers {
		routerGroup := controller.RouterGroup

		routes := controller.Routes

		httpMethods := map[string]func(string, ...gin.HandlerFunc) gin.IRoutes{
			"GET":    routerGroup.GET,
			"POST":   routerGroup.POST,
			"PUT":    routerGroup.PUT,
			"DELETE": routerGroup.DELETE,
		}

		for _, route := range routes {
			if handler, exists := httpMethods[route.Method]; exists {
				handler(route.Path, route.Controller)
			}
		}

	}

}

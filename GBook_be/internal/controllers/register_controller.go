package controllers

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method     string
	Path       string
	Controller func(c *gin.Context)
}

func RegisterController(routerGroup *gin.RouterGroup, routes []Route) {
	// Ánh xạ phương thức HTTP tới các hàm xử lý tương ứng
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

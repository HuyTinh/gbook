package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
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

type ProvideControllerParams struct {
	fx.In
	Controllers []Controller `group:"controllers"`
}

func ProvideController(controllerParams ProvideControllerParams) {

	for _, controller := range controllerParams.Controllers {
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

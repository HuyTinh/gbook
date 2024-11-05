package server

import (
	"github.com/gin-gonic/gin" // Nhập gói Gin để xây dựng RESTful API
	"go.uber.org/fx"           // Nhập gói Fx để sử dụng dependency injection

	HTTPMethod "GBook_be/internal/enums" // Nhập gói enums để sử dụng các phương thức HTTP
)

// Cấu trúc Route định nghĩa một route với phương thức HTTP, đường dẫn và controller
type Route struct {
	Method     string               // Phương thức HTTP (GET, POST, PUT, DELETE)
	Path       string               // Đường dẫn cho route
	Controller func(c *gin.Context) // Hàm xử lý cho route
}

// Cấu trúc Controller chứa nhóm router và các route của nó
type Controller struct {
	RouterGroup *gin.RouterGroup // Nhóm router cho các route
	Routes      []Route          // Danh sách các route
}

// Cấu trúc ProvideControllerParams định nghĩa các tham số đầu vào cho ProvideController
type ProvideControllerParams struct {
	fx.In
	Controllers []Controller `group:"controllers"` // Nhập danh sách controllers với tag "controllers"
}

// ProvideController cung cấp các route cho Gin Router từ danh sách controllers
func ProvideController(controllerParams ProvideControllerParams) {
	// Lặp qua từng controller được cung cấp
	for _, controller := range controllerParams.Controllers {
		routerGroup := controller.RouterGroup // Lấy nhóm router từ controller
		routes := controller.Routes           // Lấy danh sách các route từ controller

		// Bản đồ phương thức HTTP đến các hàm tương ứng của Gin
		httpMethods := map[string]func(string, ...gin.HandlerFunc) gin.IRoutes{
			HTTPMethod.GET:    routerGroup.GET,    // Phương thức GET
			HTTPMethod.POST:   routerGroup.POST,   // Phương thức POST
			HTTPMethod.PUT:    routerGroup.PUT,    // Phương thức PUT
			HTTPMethod.DELETE: routerGroup.DELETE, // Phương thức DELETE
		}

		// Lặp qua từng route để đăng ký vào router
		for _, route := range routes {
			// Kiểm tra xem phương thức HTTP có tồn tại trong bản đồ không
			if handler, exists := httpMethods[route.Method]; exists {
				handler(route.Path, route.Controller) // Đăng ký route với hàm xử lý tương ứng
			}
		}
	}
}

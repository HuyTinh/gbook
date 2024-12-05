package books

import (
	"os" // Nhập gói os để truy cập biến môi trường

	"github.com/gin-gonic/gin" // Nhập gói gin để xây dựng ứng dụng web
)

// BookRoute chứa nhóm router cho các tuyến đường liên quan đến sách.
type BookRoute struct {
	route *gin.RouterGroup // Nhóm router sử dụng Gin
}

// ProvideBookRouter cung cấp BookRoute bằng cách tạo một nhóm router dựa trên gin.Engine.
func NewBookRouter(routerGroup *gin.Engine) *BookRoute {
	return &BookRoute{
		route: routerGroup.Group(os.Getenv("BOOK_END_POINT")), // Tạo nhóm router với đường dẫn từ biến môi trường BOOK_END_POINT
	}
}

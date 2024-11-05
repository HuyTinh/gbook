package authors

import (
	"os"

	"github.com/gin-gonic/gin"
)

// AuthorRoute đại diện cho nhóm route cho các tác giả.
type AuthorRoute struct {
	route *gin.RouterGroup // Nhóm router cho tác giả, cho phép định nghĩa các endpoint liên quan đến tác giả
}

// ProvideAuthorRouter cung cấp AuthorRoute với routerGroup đã cho.
func ProvideAuthorRouter(routerGroup *gin.Engine) AuthorRoute {
	// Khởi tạo nhóm route với endpoint lấy từ biến môi trường
	return AuthorRoute{
		route: routerGroup.Group(os.Getenv("AUTHOR_END_POINT")), // Tạo một nhóm route con từ router chính
	}
}

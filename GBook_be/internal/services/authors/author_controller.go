package authors

import (
	HTTPMethod "GBook_be/internal/enums" // Nhập gói enums để sử dụng các phương thức HTTP
	"GBook_be/internal/server"           // Nhập gói server để sử dụng cấu trúc Controller và Route
)

// ProvideAuthorController cung cấp một Controller cho tác giả, kết nối dịch vụ với các route.
func NewAuthorController(service *AuthorService, authorRoute *AuthorRoute) server.Controller {
	return server.Controller{
		RouterGroup: authorRoute.route, // Nhóm router cho các route của tác giả
		Routes: []server.Route{ // Định nghĩa các route cho controller
			{
				Method:     HTTPMethod.GET,       // Phương thức HTTP cho route lấy danh sách tác giả
				Path:       "",                   // Đường dẫn của route (gốc cho nhóm route)
				Controller: service.GetAllAuthor, // Hàm xử lý cho route này
			},
			{
				Method:     HTTPMethod.POST,    // Phương thức HTTP cho route lưu tác giả mới
				Path:       "",                 // Đường dẫn của route (gốc cho nhóm route)
				Controller: service.SaveAuthor, // Hàm xử lý cho route này
			},
		},
	}
}

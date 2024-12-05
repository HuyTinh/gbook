package books

import (
	HTTPMethod "GBook_be/internal/enums" // Nhập các phương thức HTTP từ gói enums
	"GBook_be/internal/server"           // Nhập gói server để định nghĩa Controller và Route
)

// ProvideBookController thiết lập controller cho sách với các tuyến đường và dịch vụ của nó.
func NewBookController(service *BookService, bookRoute *BookRoute) server.Controller {
	return server.Controller{
		RouterGroup: bookRoute.route, // Đặt nhóm router cho controller
		Routes: []server.Route{
			{
				Method:     HTTPMethod.GET,     // Phương thức HTTP để lấy tất cả sách
				Path:       "",                 // Đường dẫn cho tuyến đường (gốc của endpoint sách)
				Controller: service.GetAllBook, // Hàm controller để xử lý yêu cầu
			},
			{
				Method:     HTTPMethod.GET,       // Phương thức HTTP để lấy tất cả sách
				Path:       "/:id",               // Đường dẫn cho tuyến đường (gốc của endpoint sách)
				Controller: service.FindBookById, // Hàm controller để xử lý yêu cầu
			},
			{
				Method:     HTTPMethod.GET,         // Phương thức HTTP để lấy tất cả sách
				Path:       "/slug/:slug",          // Đường dẫn cho tuyến đường (gốc của endpoint sách)
				Controller: service.FindBookBySlug, // Hàm controller để xử lý yêu cầu
			},
			{
				Method:     HTTPMethod.POST,  // Phương thức HTTP để lưu một sách mới
				Path:       "",               // Đường dẫn cho tuyến đường (gốc của endpoint sách)
				Controller: service.SaveBook, // Hàm controller để xử lý yêu cầu
			},
		},
	}
}

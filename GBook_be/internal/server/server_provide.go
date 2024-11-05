package server

import (
	"GBook_be/internal/database" // Nhập gói database để cung cấp kết nối cơ sở dữ liệu

	"go.uber.org/fx" // Nhập gói fx để sử dụng dependency injection
)

// ProvideServers cung cấp các server cho ứng dụng bằng cách sử dụng dependency injection.
func ProvideServers() fx.Option {
	return fx.Provide(
		ProvideServer,            // Cung cấp hàm ProvideServer để khởi tạo server
		ProvideRoutes,            // Cung cấp hàm ProvideRoutes để khởi tạo các route
		database.ProvideDatabase, // Cung cấp hàm ProvideDatabase để khởi tạo kết nối cơ sở dữ liệu
	)
}

package books

import (
	"go.uber.org/fx" // Nhập gói fx từ Uber để quản lý dependency injection
)

// ProvideBooks cung cấp các thành phần liên quan đến sách dưới dạng fx.Option.
func ProvideBooks() fx.Option {
	return fx.Provide(
		ProvideBookRepository, // Cung cấp repository cho sách
		ProvideBookService,    // Cung cấp service cho sách
		ProvideBookRouter,     // Cung cấp router cho sách
		fx.Annotate(ProvideBookController, // Cung cấp controller cho sách
			fx.ResultTags(`group:"controllers"`), // Đánh dấu kết quả trong nhóm "controllers"
		),
	)
}

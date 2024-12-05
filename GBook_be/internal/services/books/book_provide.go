package books

import (
	"go.uber.org/fx" // Nhập gói fx từ Uber để quản lý dependency injection
)

// ProvideBooks cung cấp các thành phần liên quan đến sách dưới dạng fx.Option.
func ProvideBooks() fx.Option {
	return fx.Provide(
		NewBookRepository, // Cung cấp repository cho sách
		NewBookService,    // Cung cấp service cho sách
		NewBookRouter,     // Cung cấp router cho sách
		fx.Annotate(NewBookController, // Cung cấp controller cho sách
			fx.ResultTags(`group:"controllers"`), // Đánh dấu kết quả trong nhóm "controllers"
		),
	)
}

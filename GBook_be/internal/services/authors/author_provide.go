package authors

import (
	"go.uber.org/fx" // Nhập gói fx để sử dụng dependency injection
)

// ProvideAuthors cung cấp các thành phần liên quan đến tác giả cho Fx.
func ProvideAuthors() fx.Option {
	return fx.Provide(
		NewAuthorRepository, // Cung cấp AuthorRepository
		NewAuthorService,    // Cung cấp AuthorService
		NewAuthorRouter,     // Cung cấp AuthorRouter
		fx.Annotate(NewAuthorController, // Cung cấp AuthorController với tag
			fx.ResultTags(`group:"controllers"`), // Đánh dấu để nhóm controllers
		),
	)
}

package authors

import (
	"go.uber.org/fx" // Nhập gói fx để sử dụng dependency injection
)

// ProvideAuthors cung cấp các thành phần liên quan đến tác giả cho Fx.
func ProvideAuthors() fx.Option {
	return fx.Provide(
		ProvideAuthorRepository, // Cung cấp AuthorRepository
		ProvideAuthorService,    // Cung cấp AuthorService
		ProvideAuthorRouter,     // Cung cấp AuthorRouter
		fx.Annotate(ProvideAuthorController, // Cung cấp AuthorController với tag
			fx.ResultTags(`group:"controllers"`), // Đánh dấu để nhóm controllers
		),
	)
}

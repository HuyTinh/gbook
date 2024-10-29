package books

import (
	"go.uber.org/fx"
)

func ProvideBooks() fx.Option {
	return fx.Provide(
		ProvideBookRepository,
		ProvideBookService,
		ProvideBookRouter,
		fx.Annotate(ProvideBookController,
			fx.ResultTags(`group:"controllers"`),
		),
	)
}

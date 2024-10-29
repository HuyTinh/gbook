package authors

import (
	"go.uber.org/fx"
)

func ProvideAuthors() fx.Option {
	return fx.Provide(
		ProvideAuthorRepository,
		ProvideAuthorService,
		ProvideAuthorRouter,
		fx.Annotate(ProvideAuthorController,
			fx.ResultTags(`group:"controllers"`),
		),
	)
}

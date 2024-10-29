package server

import (
	"GBook_be/internal/database"

	"go.uber.org/fx"
)

func ProvideServers() fx.Option {
	return fx.Provide(
		ProvideServer,
		ProvideRoutes,
		database.ProvideDatabase,
	)
}

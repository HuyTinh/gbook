package authors

import (
	HTTPMethod "GBook_be/internal/enums"
	"GBook_be/internal/server"
)

func ProvideAuthorController(service AuthorService, authorRoute AuthorRoute) server.Controller {
	return server.Controller{
		RouterGroup: authorRoute.route,
		Routes: []server.Route{
			{
				Method:     HTTPMethod.GET,
				Path:       "",
				Controller: service.GetAllAuthor,
			},
			{
				Method:     HTTPMethod.POST,
				Path:       "",
				Controller: service.SaveAuthor,
			},
		},
	}
}

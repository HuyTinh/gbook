package books

import (
	HTTPMethod "GBook_be/internal/enums"
	"GBook_be/internal/server"
)

func ProvideBookController(service BookService, bookRoute BookRoute) server.Controller {
	return server.Controller{
		RouterGroup: bookRoute.route,
		Routes: []server.Route{
			{
				Method:     HTTPMethod.GET,
				Path:       "",
				Controller: service.GetAllBooks,
			},
			{
				Method:     HTTPMethod.POST,
				Path:       "",
				Controller: service.SaveBook,
			},
		},
	}
}

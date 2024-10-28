package controllers

import (
	"GBook_be/internal/services"

	HTTPMethod "GBook_be/internal/enums"
)

func BookController(service *services.BookService) []Route {
	return []Route{
		{
			Method:     HTTPMethod.GET,
			Path:       "",
			Controller: service.GetAllBooks,
		},
	}
}

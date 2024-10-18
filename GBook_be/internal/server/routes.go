package server

import (
	"GBook_be/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	bookHandler := controllers.NewBookHandler()

	r.GET("/books", bookHandler.GetAllBook)

	// r.GET("/health", s.healthHandler)

	return r
}

package server

import (
	"GBook_be/internal/database"
	"GBook_be/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Initial database
	db := database.DB

	routes.RegisterBookRouter(r, db)

	return r
}

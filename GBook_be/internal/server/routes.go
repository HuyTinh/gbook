package server

import (
	"GBook_be/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	routes.RegisterBookRouter(r)

	return r
}

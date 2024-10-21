package server

import (
	"GBook_be/internal/database"
	"GBook_be/internal/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                      // Allow specific origins
		AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Cache preflight response for 12 hours
	}))

	// Initialize database
	db := database.DB

	routes.RegisterBookRouter(r, db)

	return r
}

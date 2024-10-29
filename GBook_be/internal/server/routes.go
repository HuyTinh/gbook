package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ProvideRoutes() *gin.Engine {
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

	return r
}

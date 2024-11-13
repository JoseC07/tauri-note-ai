package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"tauri-notes-ai/database"
	"tauri-notes-ai/routes"
	"time"
)

func main() {
	// Initialize Database
	database.InitDB()

	// Initialize Gin with release mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1420", "tauri://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Setup routes
	api := r.Group("/api")
	{
		routes.SetupAudioRoutes(api)
		routes.SetupNoteRoutes(api)
	}

	r.Run(":8080")
} 
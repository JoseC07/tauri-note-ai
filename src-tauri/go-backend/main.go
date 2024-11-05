package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tauri-notes-ai/database"
	"tauri-notes-ai/routes"
	"tauri-notes-ai/setup"
)

func main() {
	// POTENTIAL ISSUE: Need to handle environment configuration across platforms
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize required directories
	if err := setup.EnsureAppDirectories(); err != nil {
		log.Fatal("Failed to create app directories:", err)
	}

	// Check AI models
	if err := setup.CheckAIModels(); err != nil {
		log.Fatal("AI models not properly installed:", err)
	}

	// Initialize database
	database.InitDB()

	// Setup Gin
	r := gin.Default()

	// POTENTIAL ISSUE: Need to configure CORS properly for Tauri
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Setup routes
	routes.SetupRoutes(r)

	// Graceful shutdown handling
	// POTENTIAL ISSUE: Need to ensure all resources are properly cleaned up
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		// Cleanup code here
		os.Exit(0)
	}()

	// POTENTIAL ISSUE: Need to handle port conflicts
	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 
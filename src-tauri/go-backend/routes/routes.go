package routes

import (
	"github.com/gin-gonic/gin"
	"tauri-notes-ai/controllers"
)

func SetupRoutes(r *gin.Engine) {
	// Notes routes
	r.GET("/notes", controllers.GetNotes)
	r.POST("/notes", controllers.CreateNote)
	
	// Audio routes
	r.GET("/devices", controllers.ListAudioDevices)
	r.POST("/audio/process", controllers.ProcessAudioWithPython)
} 
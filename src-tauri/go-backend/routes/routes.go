package routes

import (
	"github.com/gin-gonic/gin"
	"tauri-notes-ai/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/notes", controllers.GetNotes)
		api.POST("/notes", controllers.CreateNote)
		api.GET("/audio-devices", controllers.ListAudioDevices)
	}
} 
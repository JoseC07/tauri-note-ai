package controllers

import (
	"github.com/gin-gonic/gin"
	"tauri-notes-ai/database"
	"tauri-notes-ai/models"
	"net/http"
	"os/exec"
	"encoding/json"
)

func GetNotes(c *gin.Context) {
	var notes []models.Note
	result := database.DB.Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
	// Parse multipart form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Access the uploaded file
	file, _, err := c.Request.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Audio file is required"})
		return
	}
	defer file.Close()

	// Save the audio file to disk (or process as needed)
	// For simplicity, we'll skip saving the file

	// Read the audio data (here we assume it's text for simplicity)
	audioData := "Simulated audio data"

	// Call the Python script
	cmd := exec.Command("python", "src-tauri/python/notes_ai_helper.py", audioData)
	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to run AI script"})
		return
	}

	// Parse the output from the AI script
	var aiResult map[string]string
	if err := json.Unmarshal(output, &aiResult); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid AI script output"})
		return
	}

	// Create a new note instance
	note := models.Note{
		Summary: aiResult["summary"],
	}

	// Save the note to the database
	if err := database.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created note as JSON
	c.JSON(http.StatusOK, note)
} 
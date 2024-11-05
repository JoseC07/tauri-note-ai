package controllers

import (
	"github.com/gin-gonic/gin"
	"tauri-notes-ai/database"
	"tauri-notes-ai/models"
)

// POTENTIAL ISSUE: Need to implement pagination for large datasets
func GetNotes(c *gin.Context) {
	var notes []models.Note
	result := database.DB.Find(&notes)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch notes"})
		return
	}
	c.JSON(200, notes)
}

// POTENTIAL ISSUE: Need input validation and sanitization
func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// POTENTIAL ISSUE: Need to handle concurrent writes
	result := database.DB.Create(&note)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create note"})
		return
	}
	
	c.JSON(201, note)
} 
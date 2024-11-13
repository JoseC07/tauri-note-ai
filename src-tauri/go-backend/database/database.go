package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tauri-notes-ai/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Note{})
} 
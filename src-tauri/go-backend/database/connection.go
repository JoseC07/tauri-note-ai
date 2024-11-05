package database

import (
	"log"
	"os"
	"path/filepath"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tauri-notes-ai/models"
)

var DB *gorm.DB

func InitDB() {
	// POTENTIAL ISSUE: Need to handle database file location across different platforms
	dbPath := filepath.Join(os.Getenv("APP_DATA_DIR"), "notes.db")
	
	// POTENTIAL ISSUE: Need proper error handling for DB connection failures
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	// Auto Migrate the schema
	// POTENTIAL ISSUE: Migrations might need more sophisticated handling for production
	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	DB = db
} 
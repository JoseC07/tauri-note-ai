package models

import (
	"time"
	"gorm.io/gorm"
)

type Note struct {
	ID        uint           `gorm:"primaryKey"`
	Content   string         `gorm:"type:text"`
	Summary   string         `gorm:"type:text"`
	AudioPath string         `gorm:"type:text"` // POTENTIAL ISSUE: Need to consider storage strategy for audio files
	Timestamp time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`     // Soft deletes
}

// POTENTIAL ISSUE: Large text fields might need different storage strategy for performance
// POTENTIAL ISSUE: Need to consider indexing strategy for search functionality 
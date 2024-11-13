package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Summary string `json:"summary"`
}
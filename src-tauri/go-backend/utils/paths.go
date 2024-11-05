package utils

import (
	"os"
	"path/filepath"
)

// GetAppDataDir returns the application data directory
func GetAppDataDir() string {
	appDir := os.Getenv("APP_DATA_DIR")
	if appDir == "" {
		userConfigDir, err := os.UserConfigDir()
		if err != nil {
			return "."
		}
		return filepath.Join(userConfigDir, "tauri-notes-ai")
	}
	return appDir
} 
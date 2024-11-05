package setup

import (
	"os"
	"path/filepath"
)

func EnsureAppDirectories() error {
	// Get user's app data directory
	appDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	// Create necessary directories
	dirs := []string{
		filepath.Join(appDir, "tauri-notes-ai", "database"),
		filepath.Join(appDir, "tauri-notes-ai", "audio"),
		filepath.Join(appDir, "tauri-notes-ai", "models"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func CheckAIModels() error {
	// Check if required AI models are present
	// TODO: Implement model verification
	return nil
} 
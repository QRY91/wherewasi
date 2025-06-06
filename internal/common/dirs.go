package common

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetDataDir returns the appropriate data directory for the current OS
// Linux/macOS: ~/.local/share/wherewasi (XDG compliant)
// Windows: %APPDATA%\wherewasi
func GetDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Fallback to current directory if we can't get home
		return filepath.Join(".", "wherewasi")
	}

	if runtime.GOOS == "windows" {
		// Windows: Use AppData/Roaming
		appData := os.Getenv("APPDATA")
		if appData != "" {
			return filepath.Join(appData, "wherewasi")
		}
		// Fallback if APPDATA not set
		return filepath.Join(homeDir, "AppData", "Roaming", "wherewasi")
	}

	// Linux/macOS: XDG compliant
	return filepath.Join(homeDir, ".local", "share", "wherewasi")
}

// GetConfigDir returns the cross-platform config directory
// Linux/macOS: ~/.config/wherewasi (XDG compliant)
// Windows: %APPDATA%/wherewasi
func GetConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Fallback to current directory if we can't get home
		return filepath.Join(".", "wherewasi")
	}

	if runtime.GOOS == "windows" {
		// Windows: Use AppData/Roaming
		appData := os.Getenv("APPDATA")
		if appData != "" {
			return filepath.Join(appData, "wherewasi")
		}
		// Fallback if APPDATA not set
		return filepath.Join(homeDir, "AppData", "Roaming", "wherewasi")
	}

	// Linux/macOS: XDG compliant
	return filepath.Join(homeDir, ".config", "wherewasi")
}

// GetDefaultDBPath returns the default database path
// Linux/macOS: ~/.local/share/wherewasi/context.sqlite
// Windows: %APPDATA%/wherewasi/context.sqlite
func GetDefaultDBPath() string {
	return filepath.Join(GetDataDir(), "context.sqlite")
}

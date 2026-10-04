package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
)

const (
	ConfigDirName  = "gacha-planner"
	ConfigFileName = "data.json"
)

// GetConfigFilePath resolves the full absolute path to ~/.config/gacha-planner/data.json
func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", ConfigDirName)
	return filepath.Join(configDir, ConfigFileName), nil
}

// LoadProfile reads the user profile from disk. If the file doesn't exist,
// it creates a default profile with 0 jades/drops and saves it.
func LoadProfile() (models.UserProfile, error) {
	filePath, err := GetConfigFilePath()
	if err != nil {
		return models.NewZeroProfile(), err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist -> Create default zero profile & save
		defaultProfile := models.NewZeroProfile()
		if saveErr := SaveProfile(defaultProfile); saveErr != nil {
			return defaultProfile, fmt.Errorf("failed to initialize default data file: %w", saveErr)
		}
		return defaultProfile, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return models.NewZeroProfile(), fmt.Errorf("failed to read data file: %w", err)
	}

	var profile models.UserProfile
	if err := json.Unmarshal(data, &profile); err != nil {
		return models.NewZeroProfile(), fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return profile, nil
}

// SaveProfile writes the user profile to disk, creating parent directories if needed.
func SaveProfile(profile models.UserProfile) error {
	filePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory %s: %w", dir, err)
	}

	data, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal profile JSON: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write data file %s: %w", filePath, err)
	}

	return nil
}

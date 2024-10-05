package configutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestEnsureExists(t *testing.T) {
	tests := []struct {
		name           string
		setupConfig    bool
		initialContent string
	}{
		{
			name:           "CreatesConfigFile",
			setupConfig:    false,
			initialContent: "",
		},
		{
			name:           "ConfigFileExists",
			setupConfig:    true,
			initialContent: "apiKeys:\n  openWeather: test-openweather-api-key\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save the original HOME environment variable
			originalHome := os.Getenv("HOME")

			// Create a temporary directory
			tempDir, err := os.MkdirTemp("", "weather-cli-test")
			if err != nil {
				t.Fatalf("Failed to create temp directory: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Set HOME to the temporary directory
			os.Setenv("HOME", tempDir)
			defer os.Setenv("HOME", originalHome)

			// Reset Viper
			viper.Reset()

			configPath := filepath.Join(tempDir, ".weather")
			configFile := filepath.Join(configPath, "config")

			if tt.setupConfig {
				// Create the config directory and file
				err = os.MkdirAll(configPath, os.ModePerm)
				if err != nil {
					t.Fatalf("Failed to create config directory: %v", err)
				}

				// Write initial content to the config file
				err = os.WriteFile(configFile, []byte(tt.initialContent), 0644)
				if err != nil {
					t.Fatalf("Failed to write to config file: %v", err)
				}
			}

			// Call EnsureExists
			EnsureExists()

			// Verify that the config file exists
			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				t.Errorf("Config file does not exist at %s", configFile)
			}

			if tt.setupConfig {
				// Read the config file to verify the content is preserved
				content, err := os.ReadFile(configFile)
				if err != nil {
					t.Fatalf("Failed to read config file: %v", err)
				}

				if string(content) != tt.initialContent {
					t.Errorf("Config file content was modified by EnsureExists")
				}
			}
		})
	}
}

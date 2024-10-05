package configutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// EnsureExists checks if the configuration file exists and initializes it if not.
func EnsureExists() {
	configPath := filepath.Join(os.Getenv("HOME"), ".weather")
	configFile := filepath.Join(configPath, "config")

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// If the config file doesn't exist, create it
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		os.MkdirAll(configPath, os.ModePerm)
		file, err := os.Create(configFile)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			os.Exit(1)
		}
		file.Close()
	}

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
}

// RemoveKey deletes a configuration key or keys from the config file.
func RemoveKey(vars ...string) error {
	cfg := viper.AllSettings()
	vals := cfg

	for _, v := range vars {
		parts := strings.Split(v, ".")
		for i, k := range parts {
			v, ok := vals[k]
			if !ok {
				// Doesn't exist no action needed
				break
			}

			switch len(parts) {
			case i + 1:
				// Last part so delete.
				delete(vals, k)
			default:
				m, ok := v.(map[string]interface{})
				if !ok {
					return fmt.Errorf("unsupported type: %T for %q", v, strings.Join(parts[0:i], "."))
				}
				vals = m
			}
		}
	}

	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err = viper.ReadConfig(bytes.NewReader(b)); err != nil {
		return err
	}

	return viper.WriteConfig()
}

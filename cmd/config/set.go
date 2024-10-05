package config

import (
	"fmt"

	"github.com/sebvaldez/weather-cli/internal/configutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// Add 'set' command with flags
	ConfigCmd.AddCommand(configSetCmd)
	configSetCmd.Flags().StringVar(&owAPIKey, "ow-api-key", "", "Set OpenWeather API key")
	configSetCmd.Flags().StringVar(&ip2locationKey, "ip2location-api-key", "", "Set ip2location API key")
	configSetCmd.Flags().StringVar(&units, "units", "", "Set units (standard, metric, imperial)")
}

var (
	owAPIKey       string
	ip2locationKey string
	units          string
)
var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration values",

	Run: func(cmd *cobra.Command, args []string) {
		// Ensure the config file exists.
		configutils.EnsureExists()

		// Update config file with values based on flags.
		if owAPIKey != "" {
			viper.Set("apiKeys.openWeather", owAPIKey)
			fmt.Println("OpenWeather api key set.")
		}
		if ip2locationKey != "" {
			viper.Set("apiKeys.ip2location", ip2locationKey)
			fmt.Println("Ip2Location api key set.")
		}
		if units != "" {
			if units != "metric" && units != "imperial" && units != "standard" {
				fmt.Println("Invalid units. Please use 'standard', 'metric', or 'imperial'.")
				return
			}
			viper.Set("units", units)
			fmt.Println("Units set to:", units)
		}

		// Save the configuration
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Error saving configuration:", err)
			return
		}

	},
}

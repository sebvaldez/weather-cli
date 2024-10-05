package config

import (
	"fmt"

	"github.com/sebvaldez/weather-cli/internal/configutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View configuration values",

	Run: func(cmd *cobra.Command, args []string) {
		// Ensure the config file exists.
		configutils.EnsureExists()

		allSettings := viper.AllSettings()

		config, err := yaml.Marshal(&allSettings)
		if err != nil {
			fmt.Println("Error when reading configurations:", err)
			return
		}

		fmt.Println(string(config))
	},
}

func init() {
	// Add 'view' command with flags
	ConfigCmd.AddCommand(configViewCmd)
}

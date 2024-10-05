package config

import (
	"fmt"

	"github.com/sebvaldez/weather-cli/internal/configutils"
	"github.com/spf13/cobra"
)

func init() {
	// Add 'unset' command
	ConfigCmd.AddCommand(configUnSetCmd)
	configUnSetCmd.Flags().BoolVar(&unsetOwAPIKey, "ow-api-key", false, "Unset OpenWeather API key")
	configUnSetCmd.Flags().BoolVar(&unsetIp2locationKey, "ip2location-api-key", false, "Unset Ip2Location API key")
	configUnSetCmd.Flags().BoolVar(&unsetUnits, "units", false, "Unset preferred units")

}

var (
	unsetOwAPIKey       bool
	unsetIp2locationKey bool
	unsetUnits          bool
)
var configUnSetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset configuration values",

	Run: func(cmd *cobra.Command, args []string) {
		configutils.EnsureExists()

		// Unset values based on flags
		if unsetOwAPIKey {
			err := configutils.RemoveKey("apiKeys.openWeather")
			if err != nil {
				fmt.Println("Error unsetting OpenWeather API key:", err)
			} else {
				fmt.Println("OpenWeather API key unset.")
			}
		}

		if unsetIp2locationKey {
			err := configutils.RemoveKey("apiKeys.ip2location")
			if err != nil {
				fmt.Println("Error unsetting ip2location API key:", err)
			} else {
				fmt.Println("ip2location API key unset.")
			}
		}

		if unsetUnits {
			err := configutils.RemoveKey("units")
			if err != nil {
				fmt.Println("Error unsetting units:", err)
			} else {
				fmt.Println("Units unset.")
			}
		}
	},
}

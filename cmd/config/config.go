package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure weather-cli",
	Long:  `Modify weather-cli configuration such as API keys, units, and more.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
}

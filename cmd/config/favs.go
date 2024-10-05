package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// Add 'favs' command
	ConfigCmd.AddCommand(FavsCmd)
	// require a --name, --state, --city flag
	FavsCmd.Flags().StringP("name", "n", "", "Name of the location")
	FavsCmd.Flags().StringP("state", "s", "", "State of the location")
	FavsCmd.Flags().StringP("city", "c", "", "City of the location")
}

var FavsCmd = &cobra.Command{
	Use:   "favs",
	Short: "Manage your favorite locations",
	Long:  `Add, remove, or list your favorite locations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("favs called")
	},
}

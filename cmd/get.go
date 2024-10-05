/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sebvaldez/weather-cli/internal/openweather"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Looks up the current weather for a given location",

	Run: func(cmd *cobra.Command, args []string) {

		weatherClient := openweather.NewClient(os.Getenv("OPENWEATHER_API_KEY"))
		res, err := weatherClient.Get()
		if err != nil {
			fmt.Printf("error getting weather: %v\n", err)
			os.Exit(1)
		}

		// Print current weather and temp
		fmt.Printf("Current temperature: %.2f°F\n", convertKelvinToFahrenheit(res.Current.Temp))

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func convertKelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sebvaldez/weather-cli/cmd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "weather",
	Short:   "A simple CLI tool to fetch weather information",
	Aliases: []string{"wctl"},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, ".weather")
	configFile := filepath.Join(configPath, "config")

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found ...
		} else {
			fmt.Println("Error reading config file:", err)
		}
	}

	rootCmd.AddCommand(config.ConfigCmd)
}

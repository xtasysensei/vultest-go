/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vultest",
	Short: "A cross-site scripting and sql injection scanner",
	Long: `A cli tool written in go that scans a
given url for cross-site scripting and sql injection vulnerability`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vultest.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

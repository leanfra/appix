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
	Use:   "cli",
	Short: "Appix CLI",
	Long:  `Appix CLI is client for managing Appix system.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string
var serverAddr string

func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "~/.appix/config.yaml", "config file (default is $HOME/.appix/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&serverAddr, "server", "s", "127.0.0.1:8080", "Appix server address. default 127.0.0.1:8080")

}

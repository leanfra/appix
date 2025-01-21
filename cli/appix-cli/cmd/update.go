/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update one or more resources",
	Long:  `Update one or more resources with the specified ID and fields.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var updateFile string
var updateOnline bool

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")
	updateCmd.PersistentFlags().StringVarP(&updateFile, "yaml", "f", "", "YAML file to update (cannot be used with --edit)")
	updateCmd.PersistentFlags().BoolVarP(&updateOnline, "edit", "e", false, "Edit online (cannot be used with --yaml)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources like apps, clusters, datacenters, environments and features",
	Long:  `Create various resources like applications, clusters, datacenters, environments and features.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var yamlFile string
var outFile string

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	createCmd.PersistentFlags().StringVarP(&yamlFile, "yaml", "f", "", "Name of the resource file to create")
	createCmd.PersistentFlags().StringVarP(&outFile, "out", "o", "", "Name of the template file to create")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

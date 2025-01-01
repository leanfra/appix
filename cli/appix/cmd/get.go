/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Appix resources",
	Long:  `Get Appix resources.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFormat(GetFormat)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

var GetFormat string
var ValidGetFormats = []string{"table", "yaml", "text"}

const DefaultPageSize = uint32(50)
const DefaultPage = uint32(1)

var GetPageSize = DefaultPageSize
var GetPage = DefaultPage

func validateFormat(format string) error {
	for _, f := range ValidGetFormats {
		if f == format {
			return nil
		}
	}
	return fmt.Errorf("invalid format %q, valid values are: table, yaml, text", format)
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&GetFormat, "format", "f", "table", "Output format. table or yaml or text")
	getCmd.PersistentFlags().Uint32VarP(&GetPageSize, "page-size", "p", GetPageSize, "Page size")
	getCmd.PersistentFlags().Uint32VarP(&GetPage, "page", "P", GetPage, "Page")

}

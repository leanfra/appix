/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

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

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "~/.appix/config.yaml", "config file")
	rootCmd.PersistentFlags().StringVarP(&serverAddr, "server", "s", "127.0.0.1:9000", "Appix server address")

}

// 辅助函数：转换 []uint 到 []uint32
func toUint32Slice(slice []uint) []uint32 {
	result := make([]uint32, len(slice))
	for i, v := range slice {
		result[i] = uint32(v)
	}
	return result
}

// 添加辅助函数
func joinUint32(ids []uint32) string {
	if len(ids) == 0 {
		return ""
	}

	strs := make([]string, len(ids))
	for i, id := range ids {
		strs[i] = fmt.Sprint(id)
	}
	return strings.Join(strs, ",")
}

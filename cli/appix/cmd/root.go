/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"gopkg.in/yaml.v2"
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

func NewConnection(withToken bool) (context.Context, *grpc.ClientConn, error) {
	ctx := context.Background()
	if withToken {
		// Read token from config file
		configPath := strings.Replace(cfgFile, "~", os.Getenv("HOME"), 1)
		existingConfig := make(map[string]string)
		if existingData, err := os.ReadFile(configPath); err == nil {
			if err := yaml.Unmarshal(existingData, &existingConfig); err != nil {
				fmt.Printf("Failed to parse config: %v\n", err)
				return nil, nil, err
			}
		}

		// Get token
		token := existingConfig["token"]
		if token == "" {
			fmt.Println("No active session found")
			return nil, nil, fmt.Errorf("no active session found")
		}

		// Create context with token
		md := metadata.New(map[string]string{
			"Authorization": "Bearer " + token,
		})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	// Create gRPC client
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		return nil, nil, err
	}
	return ctx, conn, nil
}

// toUint32Slice: []uint to  []uint32
func toUint32Slice(slice []uint) []uint32 {
	result := make([]uint32, len(slice))
	for i, v := range slice {
		result[i] = uint32(v)
	}
	return result
}

// joinUint32: []uint32 to string separated by comma
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

func findEditor() string {
	editor := os.Getenv("EDITOR")
	if editor != "" {
		return editor
	}

	// Fallback editors in order of preference
	editors := []string{"vim", "nano", "notepad", "notepad.exe"}
	for _, e := range editors {
		if _, err := exec.LookPath(e); err == nil {
			return e
		}
	}
	return ""
}

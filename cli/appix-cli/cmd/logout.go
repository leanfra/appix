/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Appix server",
	Long: `Logout from Appix server and remove local authentication token.

Example:
  appix logout   # Logout and remove local token`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, conn, err := NewConnection(true)
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewAdminClient(conn)

		// Call logout API
		// Read existing config to get user info
		configPath := strings.Replace(cfgFile, "~", os.Getenv("HOME"), 1)
		existingConfig := make(map[string]string)
		if existingData, err := os.ReadFile(configPath); err == nil {
			if err := yaml.Unmarshal(existingData, &existingConfig); err != nil {
				fmt.Printf("Failed to parse config: %v\n", err)
				return
			}
		}

		// Get user info
		var user pb.User
		if userStr := existingConfig["user"]; userStr != "" {
			if err := yaml.Unmarshal([]byte(userStr), &user); err != nil {
				fmt.Printf("Failed to parse user info: %v\n", err)
				return
			}
		}

		resp, err := client.Logout(ctx, &pb.LogoutReq{
			Id: user.Id,
		})
		if err != nil {
			fmt.Printf("Connect to server failed: %v\n", err)
			return
		}

		if resp.Code != 0 {
			fmt.Printf("Logout failed: %s\n", resp.Message)
			return
		}

		// Remove token from config file
		configPath = strings.Replace(cfgFile, "~", os.Getenv("HOME"), 1)

		// Read existing config
		existingConfig = make(map[string]string)
		if existingData, err := os.ReadFile(configPath); err == nil {
			if err := yaml.Unmarshal(existingData, &existingConfig); err != nil {
				fmt.Printf("Failed to parse config: %v\n", err)
				return
			}
		}

		// Delete token
		delete(existingConfig, "token")
		delete(existingConfig, "user")

		// Write updated config
		data, err := yaml.Marshal(existingConfig)
		if err != nil {
			fmt.Printf("Failed to marshal config: %v\n", err)
			return
		}

		if err := os.WriteFile(configPath, data, 0644); err != nil {
			fmt.Printf("Failed to write config file: %v\n", err)
			return
		}

		fmt.Println("Logout successful")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

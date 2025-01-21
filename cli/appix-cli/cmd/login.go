/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	pb "appix/api/appix/v1"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"gopkg.in/yaml.v2"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Appix server",
	Long: `Login to Appix server with your username and password.

Examples:
  appix login --username admin --password admin123   # Login with username and password
  appix login -u admin -p admin123                  # Login with short flags

The login command will store your authentication token in ~/.appix/config.yaml
which will be used for subsequent commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get username and password from flags
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		// if password is empty, prompt for password
		if password == "" {
			fmt.Print("Password: ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println("Failed to read password")
				return
			}
			password = string(bytePassword)
			fmt.Println("")
		}

		// Create gRPC client
		ctx, conn, err := NewConnection(false)
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewAdminClient(conn)

		// Call login API
		resp, err := client.Login(ctx, &pb.LoginReq{
			UserName: username,
			Password: password,
		})

		if err != nil {
			fmt.Printf("Connect to server failed: %v\n", err)
			return
		}

		if resp.Code != 0 {
			fmt.Printf("Login failed: %s\n", resp.Message)
			return
		}

		// Save token to config file
		configPath := strings.Replace(cfgFile, "~", os.Getenv("HOME"), 1)
		configDir := filepath.Dir(configPath)

		if err := os.MkdirAll(configDir, 0755); err != nil {
			fmt.Printf("Failed to create config directory: %v\n", err)
			return
		}

		// Read existing config if it exists
		existingConfig := make(map[string]string)
		if existingData, err := os.ReadFile(configPath); err == nil {
			if err := yaml.Unmarshal(existingData, &existingConfig); err != nil {
				fmt.Printf("Failed to parse existing config: %v\n", err)
				return
			}
		}

		// Update token field
		existingConfig["token"] = resp.User.Token
		// Save user info
		user, err := yaml.Marshal(resp.User)
		if err != nil {
			fmt.Printf("Failed to marshal user: %v\n", err)
			return
		}
		existingConfig["user"] = string(user)

		// Marshal updated config
		data, err := yaml.Marshal(existingConfig)
		if err != nil {
			fmt.Printf("Failed to marshal config: %v\n", err)
			return
		}

		if err := os.WriteFile(configPath, data, 0644); err != nil {
			fmt.Printf("Failed to write config file: %v\n", err)
			return
		}

		fmt.Println("Login successful")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
	loginCmd.MarkFlagRequired("username")
}

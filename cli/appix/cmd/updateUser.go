/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "appix/api/appix/v1" // 导入 admin_grpc.pb.go 中定义的 v1 包

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// updateUserCmd represents the updateUser command
var updateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Update a user",
	Long: `Update a user with the specified ID and fields.
	The user name cannot be updated.
	Must update config.yaml if you change the admin password.`,
	Aliases: []string{"user", "users", "usr"},
	Run: func(cmd *cobra.Command, args []string) {

		ctx, conn, err := NewConnection(true)
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewAdminClient(conn)

		yamlFile, _ := cmd.Flags().GetString("yaml")
		editOnline, _ := cmd.Flags().GetBool("edit")
		var users []*pb.User
		if editOnline {
			// Get existing user data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the user data
			getReq := &pb.GetUsersRequest{Id: id}
			getResp, err := client.GetUsers(ctx, getReq)
			if err != nil {
				log.Fatalf("failed to get user: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.User{getResp.Users})
			if err != nil {
				log.Fatalf("failed to marshal user: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "user-*.yaml")
			if err != nil {
				log.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write(data); err != nil {
				log.Fatalf("failed to write temp file: %v", err)
			}
			tmpfile.Close()

			editor := findEditor()
			if editor == "" {
				log.Fatal("no suitable editor found - please set EDITOR environment variable")
			}

			cmd := exec.Command(editor, tmpfile.Name())
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("failed to run editor: %v", err)
			}

			// Read updated content
			updatedData, err := os.ReadFile(tmpfile.Name())
			if err != nil {
				log.Fatalf("failed to read updated file: %v", err)
			}
			if string(updatedData) == string(data) {
				fmt.Println("No changes detected, skipping update")
				return
			}

			// Parse updated YAML
			if err := yaml.Unmarshal(updatedData, &users); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if yamlFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(yamlFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &users); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			email, _ := cmd.Flags().GetString("email")
			phone, _ := cmd.Flags().GetString("phone")
			password, _ := cmd.Flags().GetString("password")

			users = []*pb.User{
				{
					Id:       id,
					Email:    email,
					Phone:    phone,
					Password: password,
				},
			}
		}

		req := &pb.UpdateUsersRequest{
			Users: users,
		}

		reply, err := client.UpdateUsers(ctx, req) // 使用 UpdateUsers API
		if err != nil {
			log.Fatalf("failed to update user: %v", err)
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateUserCmd)

	updateUserCmd.Flags().Uint32("id", 0, "User ID to update")
	updateUserCmd.Flags().String("email", "", "New user email")
	updateUserCmd.Flags().String("phone", "", "New user phone")
	updateUserCmd.Flags().String("password", "", "New user password")

}

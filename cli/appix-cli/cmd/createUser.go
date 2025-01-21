/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"log"

	pb "appix/api/appix/v1"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Create a new user",
	Long: `Create a new user in the system.
User is identified by name and email.

Examples:
  appix create createUser --name john --email john@example.com
  appix create createUser --out-file user-template.yaml
  appix create createUser --yaml-file user.yaml`,
	Aliases: []string{"usr", "users"},
	Run:     createUser,
}

func createUser(cmd *cobra.Command, args []string) {
	// Set up a connection to the server.
	ctx, conn, err := NewConnection(true)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAdminClient(conn)

	var req *pb.CreateUsersRequest

	if outFile != "" {
		// Generate template YAML file
		users := []*pb.User{
			{
				UserName: "user-name",
				Email:    "user@example.com",
				Password: "password",
				Phone:    "1234567890",
			},
		}

		data, err := yaml.Marshal(users)
		if err != nil {
			log.Fatalf("failed to generate yaml: %v", err)
		}

		if err := os.WriteFile(outFile, data, 0644); err != nil {
			log.Fatalf("failed to write template file: %v", err)
		}

		fmt.Printf("Template file generated at: %s\n", outFile)
		return
	} else if yamlFile != "" {
		// Read from YAML file
		data, err := os.ReadFile(yamlFile)
		if err != nil {
			log.Fatalf("failed to read yaml file: %v", err)
		}

		var user []*pb.User
		if err := yaml.Unmarshal(data, &user); err != nil {
			log.Fatalf("failed to parse yaml: %v", err)
		}

		req = &pb.CreateUsersRequest{
			Users: user,
		}
	} else {
		// Create from command line flags
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		phone, _ := cmd.Flags().GetString("phone")

		req = &pb.CreateUsersRequest{
			Users: []*pb.User{
				{
					UserName: name,
					Email:    email,
					Password: password,
					Phone:    phone,
				},
			},
		}
	}

	resp, err := client.CreateUsers(ctx, req)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	if resp != nil {
		fmt.Printf("Code: %d\n", resp.Code)
		fmt.Printf("Message: %s\n", resp.Message)
		fmt.Printf("Action: %s\n", resp.Action)
	}
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().String("name", "", "Name of the user")
	createUserCmd.Flags().String("email", "", "Email of the user")
	createUserCmd.Flags().String("password", "", "Password of the user")
	createUserCmd.Flags().String("phone", "", "Phone number of the user")
}

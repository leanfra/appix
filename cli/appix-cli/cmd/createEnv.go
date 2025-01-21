/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// createEnvCmd represents the createEnv command
var createEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Create a new environment",
	Long: `Create a new environment in the system.
Environment is used to describe the purpose of the system, like production, development, staging, etc.

Examples:
  appix create env --name prod --desc "Production Environment"
  appix create env --name dev --desc "Development Environment"`,
	Aliases: []string{"envs"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewEnvsClient(conn)

		var req *pb.CreateEnvsRequest

		if outFile != "" {
			// Generate template YAML file
			env := &pb.Env{
				Name:        "env-name",
				Description: "env description",
			}
			envs := []*pb.Env{env}

			data, err := yaml.Marshal(envs)
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

			var envs []*pb.Env
			if err := yaml.Unmarshal(data, &envs); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateEnvsRequest{
				Envs: envs,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")

			req = &pb.CreateEnvsRequest{
				Envs: []*pb.Env{
					{
						Name:        name,
						Description: desc,
					},
				},
			}
		}

		resp, err := client.CreateEnvs(ctx, req)
		if err != nil {
			log.Fatalf("failed to create environments: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createEnvCmd)
	createEnvCmd.Flags().String("name", "", "Name of the environment")
	createEnvCmd.Flags().String("desc", "", "Description of the environment")
}

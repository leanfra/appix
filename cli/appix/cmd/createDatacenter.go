/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// createDatacenterCmd represents the createDatacenter command
var createDatacenterCmd = &cobra.Command{
	Use:   "datacenter",
	Short: "Create a new datacenter",
	Long: `Create a new datacenter in the system.
Datacenter is a physical location where resources are hosted.
For cloud provider, it is a region.
Datacenter contains clusters.

Examples:
  appix create datacenter --name dc1 --desc "Primary datacenter"
  appix create datacenter --name dc2 --desc "Backup datacenter"`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewDatacentersClient(conn)

		var req *pb.CreateDatacentersRequest

		if outFile != "" {
			// Generate template YAML file
			datacenter := &pb.Datacenter{
				Name:        "datacenter-name",
				Description: "datacenter description",
			}
			datacenters := []*pb.Datacenter{datacenter}

			data, err := yaml.Marshal(datacenters)
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

			var datacenters []*pb.Datacenter
			if err := yaml.Unmarshal(data, &datacenters); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateDatacentersRequest{
				Datacenters: datacenters,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")

			req = &pb.CreateDatacentersRequest{
				Datacenters: []*pb.Datacenter{
					{
						Name:        name,
						Description: desc,
					},
				},
			}
		}

		resp, err := client.CreateDatacenters(ctx, req)
		if err != nil {
			log.Fatalf("create datacenter failed: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createDatacenterCmd)
	createDatacenterCmd.Flags().String("name", "", "Name of the datacenter")
	createDatacenterCmd.Flags().String("desc", "", "Description of the datacenter")
}

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

// createClusterCmd represents the createCluster command
var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a new cluster",
	Long: `Create a new cluster in the system.
Cluster is a group of nodes that used for same product, like k8s-1, k8s-2, etc.
Cluster contains hostgroups.

Examples:
  appix create cluster --name cluster1 --desc "Production Cluster" --env prod --dc dc1
  appix create cluster --name cluster2 --desc "Development Cluster" --env dev --dc dc2`,
	Aliases: []string{"clusters", "cls"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewClustersClient(conn)

		var req *pb.CreateClustersRequest

		if outFile != "" {
			// Generate template YAML file
			cluster := &pb.Cluster{
				Name:        "cluster-name",
				Description: "cluster description",
			}
			clusters := []*pb.Cluster{cluster}

			data, err := yaml.Marshal(clusters)
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

			var clusters []*pb.Cluster
			if err := yaml.Unmarshal(data, &clusters); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateClustersRequest{
				Clusters: clusters,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")

			req = &pb.CreateClustersRequest{
				Clusters: []*pb.Cluster{
					{
						Name:        name,
						Description: desc,
					},
				},
			}
		}

		resp, err := client.CreateClusters(ctx, req)
		if err != nil {
			log.Fatalf("failed to create clusters: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createClusterCmd)
	createClusterCmd.Flags().String("name", "", "Name of the cluster")
	createClusterCmd.Flags().String("desc", "", "Description of the cluster")
}

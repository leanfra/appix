/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// updateClusterCmd represents the updateCluster command
var updateClusterCmd = &cobra.Command{
	Use:     "cluster",
	Short:   "Update cluster",
	Long:    `Update one or more clusters with the specified ID and fields.`,
	Aliases: []string{"cluster", "clusters"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewClustersClient(conn)

		var clusters []*pb.Cluster
		if updateOnline {
			// Get existing cluster data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the cluster data
			getReq := &pb.GetClustersRequest{Id: id}
			getResp, err := client.GetClusters(cmd.Context(), getReq)
			if err != nil {
				log.Fatalf("failed to get cluster: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Cluster{getResp.Cluster})
			if err != nil {
				log.Fatalf("failed to marshal cluster: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "cluster-*.yaml")
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
			if err := yaml.Unmarshal(updatedData, &clusters); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if updateFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(updateFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &clusters); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")

			clusters = []*pb.Cluster{
				{
					Id:          id,
					Name:        name,
					Description: desc,
				},
			}
		}

		req := &pb.UpdateClustersRequest{
			Clusters: clusters,
		}

		reply, err := client.UpdateClusters(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error updating cluster: %v\n", err)
			return
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateClusterCmd)

	updateClusterCmd.Flags().Uint32("id", 0, "Cluster ID to update")
	updateClusterCmd.Flags().String("name", "", "New cluster name")
	updateClusterCmd.Flags().String("desc", "", "New cluster description")
}

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

// updateFeatureCmd represents the updateFeature command
var updateFeatureCmd = &cobra.Command{
	Use:     "feature",
	Short:   "Update feature",
	Long:    `Update one or more features with the specified ID and fields.`,
	Aliases: []string{"feature", "features"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewFeaturesClient(conn)

		var features []*pb.Feature
		if updateOnline {
			// Get existing feature data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the feature data
			getReq := &pb.GetFeaturesRequest{Id: id}
			getResp, err := client.GetFeatures(cmd.Context(), getReq)
			if err != nil {
				log.Fatalf("failed to get feature: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Feature{getResp.Feature})
			if err != nil {
				log.Fatalf("failed to marshal feature: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "feature-*.yaml")
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
			if err := yaml.Unmarshal(updatedData, &features); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}
		} else if updateFile != "" {
			data, err := os.ReadFile(updateFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &features); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			name, _ := cmd.Flags().GetString("name")
			value, _ := cmd.Flags().GetString("value")
			desc, _ := cmd.Flags().GetString("desc")

			features = []*pb.Feature{
				{
					Id:          id,
					Name:        name,
					Value:       value,
					Description: desc,
				},
			}
		}

		req := &pb.UpdateFeaturesRequest{
			Features: features,
		}

		reply, err := client.UpdateFeatures(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error updating feature: %v\n", err)
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
	updateCmd.AddCommand(updateFeatureCmd)

	updateFeatureCmd.Flags().Uint32("id", 0, "Feature ID to update")
	updateFeatureCmd.Flags().String("name", "", "New feature name")
	updateFeatureCmd.Flags().String("value", "", "New feature value")
	updateFeatureCmd.Flags().String("desc", "", "New feature description")
}

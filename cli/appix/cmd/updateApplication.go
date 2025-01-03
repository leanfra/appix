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

// updateApplicationCmd represents the updateApplication command
var updateApplicationCmd = &cobra.Command{
	Use:     "application",
	Short:   "Update application",
	Long:    `Update one or more applications with the specified ID and fields.`,
	Aliases: []string{"application", "applications", "app"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewApplicationsClient(conn)

		var applications []*pb.Application
		if updateOnline {
			// Get existing application data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the application data
			getReq := &pb.GetApplicationsRequest{Id: id}
			getResp, err := client.GetApplications(cmd.Context(), getReq)
			if err != nil {
				log.Fatalf("failed to get application: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Application{getResp.App})
			if err != nil {
				log.Fatalf("failed to marshal application: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "application-*.yaml")
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
			if err := yaml.Unmarshal(updatedData, &applications); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if updateFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(updateFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &applications); err != nil {
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
			owner, _ := cmd.Flags().GetString("owner")
			isStateful, _ := cmd.Flags().GetBool("isStateful")
			productId, _ := cmd.Flags().GetUint32("productId")
			teamId, _ := cmd.Flags().GetUint32("teamId")
			featuresId, _ := cmd.Flags().GetUint32("featuresId")
			tagsId, _ := cmd.Flags().GetUint32("tagsId")
			hostgroupsId, _ := cmd.Flags().GetUint32("hostgroupsId")

			applications = []*pb.Application{
				{
					Id:           id,
					Name:         name,
					Description:  desc,
					Owner:        owner,
					IsStateful:   isStateful,
					ProductId:    productId,
					TeamId:       teamId,
					FeaturesId:   []uint32{featuresId},
					TagsId:       []uint32{tagsId},
					HostgroupsId: []uint32{hostgroupsId},
				},
			}
		}

		req := &pb.UpdateApplicationsRequest{
			Apps: applications,
		}

		reply, err := client.UpdateApplications(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error updating application: %v\n", err)
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
	updateCmd.AddCommand(updateApplicationCmd)

	updateApplicationCmd.Flags().Uint32("id", 0, "Application ID to update")
	updateApplicationCmd.Flags().String("name", "", "New application name")
	updateApplicationCmd.Flags().String("desc", "", "New application description")
	updateApplicationCmd.Flags().String("owner", "", "New application owner")
	updateApplicationCmd.Flags().Bool("isStateful", false, "New application is stateful")
	updateApplicationCmd.Flags().Uint32("productId", 0, "New application product ID")
	updateApplicationCmd.Flags().Uint32("teamId", 0, "New application team ID")
	updateApplicationCmd.Flags().Uint32("featuresId", 0, "New application features ID")
	updateApplicationCmd.Flags().Uint32("tagsId", 0, "New application tags ID")
	updateApplicationCmd.Flags().Uint32("hostgroupsId", 0, "New application hostgroups ID")
}

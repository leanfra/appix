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
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// updateApplicationCmd represents the updateApplication command
var updateApplicationCmd = &cobra.Command{
	Use:     "app",
	Short:   "Update application",
	Long:    `Update one or more applications with the specified ID and fields.`,
	Aliases: []string{"application", "applications", "app", "apps"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
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
			getResp, err := client.GetApplications(ctx, getReq)
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
			ownerId, _ := cmd.Flags().GetUint32("ownerId")
			isStateful, _ := cmd.Flags().GetBool("isStateful")
			productId, _ := cmd.Flags().GetUint32("productId")
			teamId, _ := cmd.Flags().GetUint32("teamId")
			featuresId, _ := cmd.Flags().GetUintSlice("featuresId")
			tagsId, _ := cmd.Flags().GetUintSlice("tagsId")
			hostgroupsId, _ := cmd.Flags().GetUintSlice("hostgroupsId")

			// convert featuresId, tagsId, and hostgroupsId to uint32
			var _featuresId, _tagsId, _hostgroupsId []uint32
			for _, id := range featuresId {
				_featuresId = append(_featuresId, uint32(id))
			}
			for _, id := range tagsId {
				_tagsId = append(_tagsId, uint32(id))
			}
			for _, id := range hostgroupsId {
				_hostgroupsId = append(_hostgroupsId, uint32(id))
			}

			applications = []*pb.Application{
				{
					Id:           id,
					Name:         name,
					Description:  desc,
					OwnerId:      ownerId,
					IsStateful:   isStateful,
					ProductId:    productId,
					TeamId:       teamId,
					FeaturesId:   _featuresId,
					TagsId:       _tagsId,
					HostgroupsId: _hostgroupsId,
				},
			}
		}

		req := &pb.UpdateApplicationsRequest{
			Apps: applications,
		}

		reply, err := client.UpdateApplications(ctx, req)
		if err != nil {
			log.Fatalf("failed to update application: %v", err)
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
	updateApplicationCmd.Flags().Uint32("ownerId", 0, "New application owner")
	updateApplicationCmd.Flags().Bool("is-stateful", false, "New application is stateful")
	updateApplicationCmd.Flags().Uint32("product-id", 0, "New application product ID")
	updateApplicationCmd.Flags().Uint32("team-id", 0, "New application team ID")
	updateApplicationCmd.Flags().UintSlice("features-id", []uint{}, "New application features ID")
	updateApplicationCmd.Flags().UintSlice("tags-id", []uint{}, "New application tags ID")
	updateApplicationCmd.Flags().UintSlice("hostgroups-id", []uint{}, "New application hostgroups ID")
}

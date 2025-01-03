/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "appix/api/appix/v1"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
)

// updateHostgroupCmd represents the updateHostgroup command
var updateHostgroupCmd = &cobra.Command{
	Use:   "hostgroup",
	Short: "Update hostgroup information",
	Long: `Update hostgroup information. Can update via command line flags, YAML file, or interactive editor.
	
Examples:
  # Update via command line flags
  appix update hostgroup --id 1 --name "New Name" --desc "New description"

  # Update via YAML file
  appix update hostgroup --yaml hostgroups.yaml

  # Update interactively in editor
  appix update hostgroup --id 1 --online`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewHostgroupsClient(conn)

		var hostgroups []*pb.Hostgroup
		if updateOnline {
			// Get existing hostgroup data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the hostgroup data
			getReq := &pb.GetHostgroupsRequest{Id: id}
			getResp, err := client.GetHostgroups(cmd.Context(), getReq)
			if err != nil {
				log.Fatalf("failed to get hostgroup: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Hostgroup{getResp.Hostgroup})
			if err != nil {
				log.Fatalf("failed to marshal hostgroup: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "hostgroup-*.yaml")
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

			// Parse updated YAML
			if err := yaml.Unmarshal(updatedData, &hostgroups); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if updateFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(updateFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &hostgroups); err != nil {
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
			clusterId, _ := cmd.Flags().GetUint32("cluster-id")
			datacenterId, _ := cmd.Flags().GetUint32("datacenter-id")
			envId, _ := cmd.Flags().GetUint32("env-id")
			productId, _ := cmd.Flags().GetUint32("product-id")
			teamId, _ := cmd.Flags().GetUint32("team-id")
			uintFeatures, _ := cmd.Flags().GetUintSlice("features-id")
			uintTags, _ := cmd.Flags().GetUintSlice("tags-id")
			uintShareProducts, _ := cmd.Flags().GetUintSlice("share-products-id")
			uintShareTeams, _ := cmd.Flags().GetUintSlice("share-teams-id")

			featuresId := toUint32Slice(uintFeatures)
			tagsId := toUint32Slice(uintTags)
			shareProductsId := toUint32Slice(uintShareProducts)
			shareTeamsId := toUint32Slice(uintShareTeams)

			hostgroups = []*pb.Hostgroup{
				{
					Id:              id,
					Name:            name,
					Description:     desc,
					ClusterId:       clusterId,
					DatacenterId:    datacenterId,
					EnvId:           envId,
					ProductId:       productId,
					TeamId:          teamId,
					FeaturesId:      featuresId,
					TagsId:          tagsId,
					ShareProductsId: shareProductsId,
					ShareTeamsId:    shareTeamsId,
				},
			}
		}

		req := &pb.UpdateHostgroupsRequest{
			Hostgroups: hostgroups,
		}

		reply, err := client.UpdateHostgroups(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error updating hostgroup: %v\n", err)
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
	updateCmd.AddCommand(updateHostgroupCmd)

	updateHostgroupCmd.Flags().Uint32("id", 0, "Hostgroup ID to update")
	updateHostgroupCmd.Flags().String("name", "", "New hostgroup name")
	updateHostgroupCmd.Flags().String("desc", "", "New hostgroup description")
	updateHostgroupCmd.Flags().Uint32("cluster-id", 0, "New cluster ID")
	updateHostgroupCmd.Flags().Uint32("dc-id", 0, "New datacenter ID")
	updateHostgroupCmd.Flags().Uint32("env-id", 0, "New environment ID")
	updateHostgroupCmd.Flags().Uint32("product-id", 0, "New product ID")
	updateHostgroupCmd.Flags().Uint32("team-id", 0, "New team ID")
	updateHostgroupCmd.Flags().UintSlice("features-id", []uint{}, "New feature IDs")
	updateHostgroupCmd.Flags().UintSlice("tags-id", []uint{}, "New tag IDs")
	updateHostgroupCmd.Flags().UintSlice("share-products-id", []uint{}, "New shared product IDs")
	updateHostgroupCmd.Flags().UintSlice("share-teams-id", []uint{}, "New shared team IDs")
}

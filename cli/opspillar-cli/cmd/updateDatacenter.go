/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "opspillar/api/opspillar/v1"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// updateDatacenterCmd represents the updateDatacenter command
var updateDatacenterCmd = &cobra.Command{
	Use:     "datacenter",
	Short:   "Update datacenter",
	Long:    `Update one or more datacenters with the specified ID and fields.`,
	Aliases: []string{"datacenters", "dc"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewDatacentersClient(conn)

		var datacenters []*pb.Datacenter
		if updateOnline {
			// Get existing datacenter data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the datacenter data
			getReq := &pb.GetDatacentersRequest{Id: id}
			getResp, err := client.GetDatacenters(ctx, getReq)
			if err != nil {
				log.Fatalf("failed to get datacenter: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Datacenter{getResp.Datacenter})
			if err != nil {
				log.Fatalf("failed to marshal datacenter: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "datacenter-*.yaml")
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
			if err := yaml.Unmarshal(updatedData, &datacenters); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if updateFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(updateFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &datacenters); err != nil {
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

			datacenters = []*pb.Datacenter{
				{
					Id:          id,
					Name:        name,
					Description: desc,
				},
			}
		}

		req := &pb.UpdateDatacentersRequest{
			Datacenters: datacenters,
		}

		reply, err := client.UpdateDatacenters(ctx, req)
		if err != nil {
			log.Fatalf("failed to update datacenter: %v", err)
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateDatacenterCmd)

	updateDatacenterCmd.Flags().Uint32("id", 0, "Datacenter ID to update")
	updateDatacenterCmd.Flags().String("name", "", "New datacenter name")
	updateDatacenterCmd.Flags().String("desc", "", "New datacenter description")
}

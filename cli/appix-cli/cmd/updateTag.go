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
	"gopkg.in/yaml.v2"
)

// updateTagCmd represents the updateTag command
var updateTagCmd = &cobra.Command{
	Use:     "tag",
	Short:   "Update tag",
	Long:    `Update one or more tags with the specified ID and fields.`,
	Aliases: []string{"tag", "tags"},
	Run: func(cmd *cobra.Command, args []string) {

		ctx, conn, err := NewConnection(true)
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewTagsClient(conn)

		yamlFile, _ := cmd.Flags().GetString("yaml")
		editOnline, _ := cmd.Flags().GetBool("edit")
		var tags []*pb.Tag
		if editOnline {
			// Get existing tag data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the tag data
			getReq := &pb.GetTagsRequest{Id: id}
			getResp, err := client.GetTags(ctx, getReq)
			if err != nil {
				log.Fatalf("failed to get tag: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Tag{getResp.Tag})
			if err != nil {
				log.Fatalf("failed to marshal tag: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "tag-*.yaml")
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
			if err := yaml.Unmarshal(updatedData, &tags); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if yamlFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(yamlFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &tags); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			key, _ := cmd.Flags().GetString("key")
			value, _ := cmd.Flags().GetString("value")
			desc, _ := cmd.Flags().GetString("desc")

			tags = []*pb.Tag{
				{
					Id:          id,
					Key:         key,
					Value:       value,
					Description: desc,
				},
			}
		}

		req := &pb.UpdateTagsRequest{
			Tags: tags,
		}

		reply, err := client.UpdateTags(ctx, req)
		if err != nil {
			log.Fatalf("failed to update tag: %v", err)
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateTagCmd)

	updateTagCmd.Flags().Uint32("id", 0, "Tag ID to update")
	updateTagCmd.Flags().String("key", "", "New tag key")
	updateTagCmd.Flags().String("value", "", "New tag value")
	updateTagCmd.Flags().String("desc", "", "New tag description")
}

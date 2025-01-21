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

// createTagCmd represents the createTag command
var createTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a new tag",
	Long: `Create a new tag in the system.
Tag is a key-value pair that used for labeling resources.

Examples:
  appix create tag --key runtime --value "Production"
  appix create tag --key runtime --value "Development"`,
	Aliases: []string{"tags"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewTagsClient(conn)

		var req *pb.CreateTagsRequest

		if outFile != "" {
			// Generate template YAML file
			tag := &pb.Tag{
				Key:   "tag-key",
				Value: "tag-value",
			}

			tags := []*pb.Tag{tag}

			data, err := yaml.Marshal(tags)
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

			var tags []*pb.Tag
			if err := yaml.Unmarshal(data, &tags); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateTagsRequest{
				Tags: tags,
			}
		} else {
			// Create from command line flags
			key, _ := cmd.Flags().GetString("key")
			value, _ := cmd.Flags().GetString("value")
			desc, _ := cmd.Flags().GetString("desc")
			req = &pb.CreateTagsRequest{
				Tags: []*pb.Tag{
					{
						Key:         key,
						Value:       value,
						Description: desc,
					},
				},
			}
		}

		resp, err := client.CreateTags(ctx, req)
		if err != nil {
			log.Fatalf("failed to create tags: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createTagCmd)
	createTagCmd.Flags().String("key", "", "Key of the tag")
	createTagCmd.Flags().String("value", "", "Value of the tag")
	createTagCmd.Flags().String("desc", "", "Description of the tag")
}

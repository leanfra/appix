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

	pb "opspillar/api/opspillar/v1"
)

// createApplicationCmd represents the createApplication command
var createApplicationCmd = &cobra.Command{
	Use:   "app",
	Short: "Create a new application",
	Long: `Create a new application in the system.
Application is a software that runs on hosts.
Application belongs to a product and team.

Examples:
  opspillar create application --name web-app --desc "Web Application" --product 1 --team 1
  opspillar create application --name api-service --desc "API Service" --product 2 --team 1`,
	Aliases: []string{"application", "applications", "apps"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewApplicationsClient(conn)

		var req *pb.CreateApplicationsRequest

		if outFile != "" {
			// Generate template YAML file
			app := &pb.Application{
				Name:         "app-name",
				Description:  "app description",
				OwnerId:      1,
				IsStateful:   false,
				ProductId:    1,
				TeamId:       1,
				FeaturesId:   []uint32{},
				TagsId:       []uint32{},
				HostgroupsId: []uint32{},
			}
			apps := []*pb.Application{app}

			data, err := yaml.Marshal(apps)
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

			var apps []*pb.Application
			if err := yaml.Unmarshal(data, &apps); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateApplicationsRequest{
				Apps: apps,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")
			owner_id, _ := cmd.Flags().GetUint32("owner-id")
			isStateful, _ := cmd.Flags().GetBool("is-stateful")
			productId, _ := cmd.Flags().GetUint32("product-id")
			teamId, _ := cmd.Flags().GetUint32("team-id")
			uintFeatures, _ := cmd.Flags().GetUintSlice("features-id")
			uintTags, _ := cmd.Flags().GetUintSlice("tags-id")
			uintHostgroups, _ := cmd.Flags().GetUintSlice("hostgroups-id")

			featuresId := toUint32Slice(uintFeatures)
			tagsId := toUint32Slice(uintTags)
			hostgroupsId := toUint32Slice(uintHostgroups)

			req = &pb.CreateApplicationsRequest{
				Apps: []*pb.Application{
					{
						Name:         name,
						Description:  desc,
						OwnerId:      owner_id,
						IsStateful:   isStateful,
						ProductId:    productId,
						TeamId:       teamId,
						FeaturesId:   featuresId,
						TagsId:       tagsId,
						HostgroupsId: hostgroupsId,
					},
				},
			}
		}

		resp, err := client.CreateApplications(ctx, req)
		if err != nil {
			log.Fatalf("failed to create applications: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createApplicationCmd)
	createApplicationCmd.Flags().String("name", "", "Name of the application")
	createApplicationCmd.Flags().String("desc", "", "Description of the application")
	createApplicationCmd.Flags().Uint32("owner-id", 0, "user id of the owner of the application")
	createApplicationCmd.Flags().Bool("is-stateful", false, "Whether the application is stateful")
	createApplicationCmd.Flags().Uint32("product-id", 0, "ID of the product this application belongs to")
	createApplicationCmd.Flags().Uint32("team-id", 0, "ID of the team this application belongs to")
	createApplicationCmd.Flags().UintSlice("features-id", []uint{}, "IDs of features this application requires")
	createApplicationCmd.Flags().UintSlice("tags-id", []uint{}, "IDs of tags for this application")
	createApplicationCmd.Flags().UintSlice("hostgroups-id", []uint{}, "IDs of hostgroups for this application")
}

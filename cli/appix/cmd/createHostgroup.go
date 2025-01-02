/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// createHostgroupCmd represents the createHostgroup command
var createHostgroupCmd = &cobra.Command{
	Use:   "hostgroup",
	Short: "Create a new hostgroup",
	Long: `Create a new hostgroup in the system.
Hostgroup is a group of hosts that used for same purpose, like web servers, database servers, etc.
Hostgroup belongs to a cluster.

Examples:
  appix create hostgroup --name web-servers --desc "Web Servers" --cluster cluster1
  appix create hostgroup --name db-servers --desc "Database Servers" --cluster cluster1`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewHostgroupsClient(conn)

		var req *pb.CreateHostgroupsRequest

		if outFile != "" {
			// Generate template YAML file
			hostgroup := &pb.Hostgroup{
				Name:            "hostgroup-name",
				Description:     "hostgroup description",
				ClusterId:       1,
				TeamId:          1,
				ProductId:       1,
				EnvId:           1,
				DatacenterId:    1,
				FeaturesId:      []uint32{},
				TagsId:          []uint32{},
				ShareProductsId: []uint32{},
				ShareTeamsId:    []uint32{},
			}
			hostgroups := []*pb.Hostgroup{hostgroup}

			data, err := yaml.Marshal(hostgroups)
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

			var hostgroups []*pb.Hostgroup
			if err := yaml.Unmarshal(data, &hostgroups); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateHostgroupsRequest{
				Hostgroups: hostgroups,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")
			clusterId, _ := cmd.Flags().GetUint32("cluster")
			teamId, _ := cmd.Flags().GetUint32("team")
			productId, _ := cmd.Flags().GetUint32("product")
			envId, _ := cmd.Flags().GetUint32("env")
			dcId, _ := cmd.Flags().GetUint32("dc")
			featureId, _ := cmd.Flags().GetUint32("feature")
			tagId, _ := cmd.Flags().GetUint32("tag")
			shareProductId, _ := cmd.Flags().GetUint32("share-product")
			shareTeamId, _ := cmd.Flags().GetUint32("share-team")

			req = &pb.CreateHostgroupsRequest{
				Hostgroups: []*pb.Hostgroup{
					{
						Name:            name,
						Description:     desc,
						ClusterId:       clusterId,
						TeamId:          teamId,
						ProductId:       productId,
						EnvId:           envId,
						DatacenterId:    dcId,
						FeaturesId:      []uint32{featureId},
						TagsId:          []uint32{tagId},
						ShareProductsId: []uint32{shareProductId},
						ShareTeamsId:    []uint32{shareTeamId},
					},
				},
			}
		}

		resp, err := client.CreateHostgroups(ctx, req)
		if err != nil {
			log.Fatalf("create hostgroup failed: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createHostgroupCmd)
	createHostgroupCmd.Flags().String("name", "", "Name of the hostgroup")
	createHostgroupCmd.Flags().String("desc", "", "Description of the hostgroup")
	createHostgroupCmd.Flags().Uint32("cluster", 0, "ID of the cluster this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("team", 0, "ID of the team this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("product", 0, "ID of the product this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("env", 0, "ID of the environment this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("dc", 0, "ID of the datacenter this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("feature", 0, "ID of the feature this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("tag", 0, "ID of the tag this hostgroup belongs to")
	createHostgroupCmd.Flags().Uint32("share-product", 0, "ID of the product this hostgroup shares with")
	createHostgroupCmd.Flags().Uint32("share-team", 0, "ID of the team this hostgroup shares with")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	pb "opspillar/api/opspillar/v1"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// matchHostgroupCmd represents the matchHostgroup command
var matchHostgroupCmd = &cobra.Command{
	Use:   "hostgroup",
	Short: "Match hostgroups by features, product and team",
	Long: `Match hostgroups by features, product and team.

Examples:
  opspillar match hostgroup --features 1,2 --product 1 --team 1`,
	Aliases: []string{"hg", "hostgroups", "hgs"},
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		features, _ := cmd.Flags().GetUintSlice("features")
		product, _ := cmd.Flags().GetUint("product")
		team, _ := cmd.Flags().GetUint("team")

		// Connect to gRPC server
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewApplicationsClient(conn)

		// Call API
		resp, err := c.MatchAppHostgroups(ctx, &pb.MatchAppHostgroupsRequest{
			FeaturesId: toUint32Slice(features),
			ProductId:  uint32(product),
			TeamId:     uint32(team),
		})
		if err != nil {
			log.Fatalf("could not match hostgroups: %v", err)
		}

		// Print results
		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "yaml":
			output := map[string]interface{}{
				"hostgroups": resp.HostgroupsId,
			}
			yamlData, err := yaml.Marshal(output)
			if err != nil {
				log.Fatalf("error formatting yaml: %v", err)
			}
			fmt.Println(string(yamlData))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Hostgroup ID"})
			for _, id := range resp.HostgroupsId {
				table.Append([]string{fmt.Sprintf("%d", id)})
			}
			table.Render()
		default:
			fmt.Printf("Matched hostgroup IDs: %v\n", resp.HostgroupsId)
		}
	},
}

func init() {
	matchCmd.AddCommand(matchHostgroupCmd)

	// Add flags
	matchHostgroupCmd.Flags().UintSliceP("features", "u", nil, "Feature IDs to match")
	matchHostgroupCmd.Flags().UintP("product", "p", 0, "Product ID to match")
	matchHostgroupCmd.Flags().UintP("team", "t", 0, "Team ID to match")
	matchHostgroupCmd.Flags().StringP("format", "f", "table", "Output format. table or yaml or text")
	// Mark required flags
	matchHostgroupCmd.MarkFlagRequired("features")
	matchHostgroupCmd.MarkFlagRequired("product")
	matchHostgroupCmd.MarkFlagRequired("team")
}

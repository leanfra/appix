package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "opspillar/api/opspillar/v1"
)

var getFeatureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Get features",
	Long: `Get features resources from the system.

Examples:
  opspillar get feature                              # List all
  opspillar get feature --names feat1,feat2          # Filter by names
  opspillar get feature --kvs key1=val1              # Filter by key-value
  opspillar get feature --ids 1,2,3                 # Filter by IDs
  opspillar get feature --names feat1 --kvs key=val  # Combined filters`,
	Aliases: []string{"feat", "features"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewFeaturesClient(conn)

		page := GetPage
		pageSize := GetPageSize
		names, _ := cmd.Flags().GetStringSlice("names")
		kvs, _ := cmd.Flags().GetStringSlice("kvs")
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		ids := make([]uint32, len(uintIds))
		for i, id := range uintIds {
			ids[i] = uint32(id)
		}

		var allFeatures []*pb.Feature
		for {
			req := &pb.ListFeaturesRequest{
				Page:     page,
				PageSize: pageSize,
				Names:    names,
				Kvs:      kvs,
				Ids:      ids,
			}

			resp, err := client.ListFeatures(ctx, req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if resp.Code != 0 {
				fmt.Printf("Response details:\n")
				fmt.Printf("  Message: %s\n", resp.Message)
				fmt.Printf("  Code: %d\n", resp.Code)
				fmt.Printf("  Action: %s\n", resp.Action)
				return
			}

			// 添加当前页的特性到结果集
			allFeatures = append(allFeatures, resp.Features...)

			// 如果返回的特性数量小于页大小，说明已经是最后一页
			if len(resp.Features) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allFeatures)
			if err != nil {
				log.Fatalf("serialize yaml failed: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Value", "Description"})
			for _, feature := range allFeatures {
				table.Append([]string{
					fmt.Sprintf("%d", feature.Id),
					feature.Name,
					feature.Value,
					feature.Description,
				})
			}
			table.Render()
		case "text":
			if len(allFeatures) == 0 {
				fmt.Println("No features found")
				return
			}
			for _, feature := range allFeatures {
				fmt.Printf("ID: %d \t Name: %s \t Value: %s \t Description: %s\n", feature.Id, feature.Name, feature.Value, feature.Description)
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getFeatureCmd)

	getFeatureCmd.Flags().StringSlice("names", []string{}, "Filter by feature names")

	getFeatureCmd.Flags().StringSlice("kvs", []string{}, "Filter by key-value pairs")

	getFeatureCmd.Flags().UintSlice("ids", []uint{}, "Filter by feature IDs")
}

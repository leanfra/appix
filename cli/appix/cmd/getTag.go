/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// getTagCmd represents the getTag command
var getTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Get tags",
	Long: `Get tags resources from the system.

This command allows you to retrieve and list tags with various filtering options:
  - Filter by tag keys (-K, --keys)
  - Filter by key-value pairs (-V, --kvs)
  - Filter by tag IDs (-I, --ids)

Output formats available:
  - table: Displays results in a formatted table
  - yaml: Outputs in YAML format
  - text: Simple text output format

Examples:
  appix get tag                     # List all tags
  appix get tag -K env,project      # Filter tags by keys
  appix get tag -V env=prod         # Filter tags by key-value pair
  appix get tag -I 1,2,3            # Filter tags by IDs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 建立 gRPC 连接
		ctx := context.Background()
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("连接失败: %v", err)
		}
		defer conn.Close()

		client := pb.NewTagsClient(conn)

		page := GetPage
		pageSize := GetPageSize
		keys, _ := cmd.Flags().GetStringSlice("keys")
		kvs, _ := cmd.Flags().GetStringSlice("kvs")
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		ids := make([]uint32, len(uintIds))
		for i, id := range uintIds {
			ids[i] = uint32(id)
		}

		var allTags []*pb.Tag
		for {
			req := &pb.ListTagsRequest{
				Filter: &pb.ListTagsFilter{
					Page:     page,
					PageSize: pageSize,
					Keys:     keys,
					Kvs:      kvs,
					Ids:      ids,
				},
			}

			resp, err := client.ListTags(ctx, req)
			if err != nil {
				log.Fatalf("获取标签失败: %v", err)
			}

			// 添加当前页的标签到结果集
			allTags = append(allTags, resp.Tags...)

			// 如果返回的标签数量小于页大小，说明已经是最后一页
			if len(resp.Tags) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allTags)
			if err != nil {
				log.Fatalf("序列化YAML失败: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Key", "Value"})
			for _, tag := range allTags {
				table.Append([]string{
					fmt.Sprintf("%d", tag.Id),
					tag.Key,
					tag.Value,
				})
			}
			table.Render()
		case "text":
			if len(allTags) == 0 {
				fmt.Println("No tags found")
				return nil
			}
			for _, tag := range allTags {
				fmt.Printf("ID: %d \t Key: %s \t Value: %s\n", tag.Id, tag.Key, tag.Value)
			}
		default: // text format
			fmt.Println("unknown format")
		}
		return nil
	},
}

func init() {
	getCmd.AddCommand(getTagCmd)

	getTagCmd.Flags().StringSliceP("keys", "K", []string{}, "Filter by key names, can specify multiple")
	getTagCmd.Flags().StringSliceP("kvs", "V", []string{}, "Filter by key-value pairs, can specify multiple")
	getTagCmd.Flags().UintSliceP("ids", "I", []uint{}, "Filter by IDs, can specify multiple")
}

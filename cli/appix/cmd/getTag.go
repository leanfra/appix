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

Examples:
  appix get tag                              # List all
  appix get tag --keys env,project           # Filter by keys
  appix get tag --kvs env=prod               # Filter by key-value
  appix get tag --ids 1,2,3                 # Filter by IDs
  appix get tag --keys env --format yaml     # Custom format`,
	Run: func(cmd *cobra.Command, args []string) {
		// 建立 gRPC 连接
		ctx := context.Background()
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
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
				log.Fatalf("serialize yaml failed: %v", err)
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
				return
			}
			for _, tag := range allTags {
				fmt.Printf("ID: %d \t Key: %s \t Value: %s\n", tag.Id, tag.Key, tag.Value)
			}
		default: // text format
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getTagCmd)

	getTagCmd.Flags().StringSlice("keys", []string{}, "Filter by key names, can specify multiple")

	getTagCmd.Flags().StringSlice("kvs", []string{}, "Filter by key-value pairs, can specify multiple")

	getTagCmd.Flags().UintSlice("ids", []uint{}, "Filter by IDs, can specify multiple")
}

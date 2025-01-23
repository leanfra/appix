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

var getDatacenterCmd = &cobra.Command{
	Use:   "datacenter",
	Short: "Get datacenters",
	Long: `Get datacenters resources from the system.

Examples:
  opspillar get datacenter                              # List all
  opspillar get datacenter --names dc1,dc2              # Filter by names
  opspillar get datacenter --ids 1,2,3                 # Filter by IDs
  opspillar get datacenter --page 1 --page-size 10     # With pagination
  opspillar get datacenter --names dc1 --format yaml   # Custom format`,
	Aliases: []string{"dc", "datacenters"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewDatacentersClient(conn)

		page := GetPage
		pageSize := GetPageSize
		names, _ := cmd.Flags().GetStringSlice("names")
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		ids := make([]uint32, len(uintIds))
		for i, id := range uintIds {
			ids[i] = uint32(id)
		}

		var allDatacenters []*pb.Datacenter
		for {
			req := &pb.ListDatacentersRequest{
				Page:     page,
				PageSize: pageSize,
				Names:    names,
				Ids:      ids,
			}

			resp, err := client.ListDatacenters(ctx, req)
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

			// 添加当前页的数据中心到结果集
			allDatacenters = append(allDatacenters, resp.Datacenters...)

			// 如果返回的数据中心数量小于页大小，说明已经是最后一页
			if len(resp.Datacenters) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allDatacenters)
			if err != nil {
				log.Fatalf("序列化YAML失败: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Description"})
			for _, dc := range allDatacenters {
				table.Append([]string{
					fmt.Sprintf("%d", dc.Id),
					dc.Name,
					dc.Description,
				})
			}
			table.Render()
		case "text":
			if len(allDatacenters) == 0 {
				fmt.Println("No datacenters found")
				return
			}
			for _, dc := range allDatacenters {
				fmt.Printf("ID: %d \t Name: %s \t Description: %s\n",
					dc.Id, dc.Name, dc.Description)
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getDatacenterCmd)

	getDatacenterCmd.Flags().StringSlice("names", []string{}, "Filter by datacenter names")

	getDatacenterCmd.Flags().UintSlice("ids", []uint{}, "Filter by datacenter IDs")
}

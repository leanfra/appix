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

var getEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Get environments",
	Long: `Get environments resources from the system.

Examples:
  opspillar get env                              # List all
  opspillar get env --names dev,prod             # Filter by names
  opspillar get env --ids 1,2,3                 # Filter by IDs
  opspillar get env --page 1 --page-size 10     # With pagination
  opspillar get env --names prod --format yaml   # Custom format`,
	Aliases: []string{"envs"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewEnvsClient(conn)

		page := GetPage
		pageSize := GetPageSize
		names, _ := cmd.Flags().GetStringSlice("names")
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		ids := make([]uint32, len(uintIds))
		for i, id := range uintIds {
			ids[i] = uint32(id)
		}

		var allEnvs []*pb.Env
		for {
			req := &pb.ListEnvsRequest{
				Page:     page,
				PageSize: pageSize,
				Names:    names,
				Ids:      ids,
			}

			resp, err := client.ListEnvs(ctx, req)
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

			// 添加当前页的环境到结果集
			allEnvs = append(allEnvs, resp.Envs...)

			// 如果返回的环境数量小于页大小，说明已经是最后一页
			if len(resp.Envs) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allEnvs)
			if err != nil {
				log.Fatalf("序列化YAML失败: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Description"})
			for _, env := range allEnvs {
				table.Append([]string{
					fmt.Sprintf("%d", env.Id),
					env.Name,
					env.Description,
				})
			}
			table.Render()
		case "text":
			if len(allEnvs) == 0 {
				fmt.Println("No environments found")
				return
			}
			for _, env := range allEnvs {
				fmt.Printf("ID: %d \t Name: %s \t Description: %s\n",
					env.Id, env.Name, env.Description)
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getEnvCmd)

	// 添加基于 ListEnvsFilter 的参数
	getEnvCmd.Flags().StringSlice("names", []string{}, "Filter by environment names")

	getEnvCmd.Flags().UintSlice("ids", []uint{}, "Filter by environment IDs")
}

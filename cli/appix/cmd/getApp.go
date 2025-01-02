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

var getAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Get applications",
	Long: `Get applications resources from the system.

Examples:
  appix get app                                    # List all
  appix get app --names app1,app2                  # Filter by names
  appix get app --ids 1,2,3                       # Filter by IDs
  appix get app --is-stateful true                # Filter stateful
  appix get app --page 1 --page-size 10           # With pagination
  appix get app --names web --clusters 1 --format yaml   # Combined filters`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewApplicationsClient(conn)

		page := GetPage
		pageSize := GetPageSize
		names, _ := cmd.Flags().GetStringSlice("names")
		isStateful, _ := cmd.Flags().GetString("is-stateful")

		// 获取并转换所有uint切片参数
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		uintProducts, _ := cmd.Flags().GetUintSlice("products")
		uintTeams, _ := cmd.Flags().GetUintSlice("teams")
		uintFeatures, _ := cmd.Flags().GetUintSlice("features")
		uintTags, _ := cmd.Flags().GetUintSlice("tags")
		uintHostgroups, _ := cmd.Flags().GetUintSlice("hostgroups")

		// 转换为uint32
		ids := toUint32Slice(uintIds)
		productsId := toUint32Slice(uintProducts)
		teamsId := toUint32Slice(uintTeams)
		featuresId := toUint32Slice(uintFeatures)
		tagsId := toUint32Slice(uintTags)
		hostgroupsId := toUint32Slice(uintHostgroups)

		var allApps []*pb.Application
		for {
			req := &pb.ListApplicationsRequest{
				Filter: &pb.ListApplicationsFilter{
					Page:         page,
					PageSize:     pageSize,
					Names:        names,
					IsStateful:   isStateful,
					Ids:          ids,
					ProductsId:   productsId,
					TeamsId:      teamsId,
					FeaturesId:   featuresId,
					TagsId:       tagsId,
					HostgroupsId: hostgroupsId,
				},
			}

			resp, err := client.ListApplications(ctx, req)
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
			allApps = append(allApps, resp.Apps...)

			if len(resp.Apps) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allApps)
			if err != nil {
				log.Fatalf("failed to generate yaml: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{
				"ID", "Name", "Description", "Owner", "Stateful",
				"Product", "Team",
				"Features", "Tags", "Hostgroups",
			})

			for _, app := range allApps {
				table.Append([]string{
					fmt.Sprint(app.Id),
					app.Name,
					app.Description,
					app.Owner,
					fmt.Sprint(app.IsStateful),
					fmt.Sprint(app.ProductId),
					fmt.Sprint(app.TeamId),
					joinUint32(app.FeaturesId),
					joinUint32(app.TagsId),
					joinUint32(app.HostgroupsId),
				})
			}
			table.Render()

		case "text":
			if len(allApps) == 0 {
				fmt.Println("No applications found")
				return
			}
			for _, app := range allApps {
				fmt.Printf("ID: %d\nName: %s\nDescription: %s\nOwner: %s\n"+
					"Stateful: %v\nProduct ID: %d\nTeam ID: %d\n"+
					"Feature IDs: %s\nTag IDs: %s\nHostgroup IDs: %s\n\n",
					app.Id, app.Name, app.Description, app.Owner,
					app.IsStateful, app.ProductId, app.TeamId,
					joinUint32(app.FeaturesId),
					joinUint32(app.TagsId),
					joinUint32(app.HostgroupsId))
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getAppCmd)

	// 只使用长格式标志
	getAppCmd.Flags().StringSlice("names", []string{}, "Filter by application names")
	getAppCmd.Flags().UintSlice("ids", []uint{}, "Filter by application IDs")
	getAppCmd.Flags().String("is-stateful", "", "Filter by stateful flag. Not set or set Empty for all. (<Empty>/true/false)")
	getAppCmd.Flags().UintSlice("products", []uint{}, "Filter by product IDs")
	getAppCmd.Flags().UintSlice("teams", []uint{}, "Filter by team IDs")
	getAppCmd.Flags().UintSlice("features", []uint{}, "Filter by feature IDs")
	getAppCmd.Flags().UintSlice("tags", []uint{}, "Filter by tag IDs")
	getAppCmd.Flags().UintSlice("hostgroups", []uint{}, "Filter by hostgroup IDs")
}

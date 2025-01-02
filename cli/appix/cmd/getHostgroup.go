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

var getHostgroupCmd = &cobra.Command{
	Use:   "hostgroup",
	Short: "Get hostgroups",
	Long: `Get hostgroups resources from the system.

Examples:
  appix get hostgroup                                # List all
  appix get hostgroup --names web,api                # Filter by names
  appix get hostgroup --ids 1,2,3                   # Filter by IDs
  appix get hostgroup --clusters 1 --teams 2        # Filter by cluster and team
  appix get hostgroup --products 1 --format yaml    # Custom format`,
	Run: func(cmd *cobra.Command, args []string) {
		page := GetPage
		pageSize := GetPageSize

		// 获取所有过滤参数
		names, _ := cmd.Flags().GetStringSlice("names")
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		clusters, _ := cmd.Flags().GetUintSlice("clusters")
		datacenters, _ := cmd.Flags().GetUintSlice("datacenters")
		envs, _ := cmd.Flags().GetUintSlice("envs")
		products, _ := cmd.Flags().GetUintSlice("products")
		teams, _ := cmd.Flags().GetUintSlice("teams")
		features, _ := cmd.Flags().GetUintSlice("features")
		tags, _ := cmd.Flags().GetUintSlice("tags")
		shareProducts, _ := cmd.Flags().GetUintSlice("share-products")
		shareTeams, _ := cmd.Flags().GetUintSlice("share-teams")

		// 转换所有 uint 切片到 uint32
		ids := toUint32Slice(uintIds)
		clustersIds := toUint32Slice(clusters)
		datacentersIds := toUint32Slice(datacenters)
		envsIds := toUint32Slice(envs)
		productsIds := toUint32Slice(products)
		teamsIds := toUint32Slice(teams)
		featuresIds := toUint32Slice(features)
		tagsIds := toUint32Slice(tags)
		shareProductsIds := toUint32Slice(shareProducts)
		shareTeamsIds := toUint32Slice(shareTeams)

		var allHostgroups []*pb.Hostgroup
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewHostgroupsClient(conn)

		for {
			req := &pb.ListHostgroupsRequest{
				Filter: &pb.ListHostgroupsFilter{
					Page:            page,
					PageSize:        pageSize,
					Names:           names,
					Ids:             ids,
					ClustersId:      clustersIds,
					DatacentersId:   datacentersIds,
					EnvsId:          envsIds,
					ProductsId:      productsIds,
					TeamsId:         teamsIds,
					FeaturesId:      featuresIds,
					TagsId:          tagsIds,
					ShareProductsId: shareProductsIds,
					ShareTeamsId:    shareTeamsIds,
				},
			}

			ctx := context.Background()
			resp, err := client.ListHostgroups(ctx, req)
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

			allHostgroups = append(allHostgroups, resp.Hostgroups...)

			if len(resp.Hostgroups) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allHostgroups)
			if err != nil {
				log.Fatalf("serialize yaml failed: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Description", "Cluster", "DC", "Env", "Product", "Team", "Features", "Tags", "ShareProducts", "ShareTeams"})
			for _, hg := range allHostgroups {
				table.Append([]string{
					fmt.Sprint(hg.Id),
					hg.Name,
					hg.Description,
					fmt.Sprint(hg.ClusterId),
					fmt.Sprint(hg.DatacenterId),
					fmt.Sprint(hg.EnvId),
					fmt.Sprint(hg.ProductId),
					fmt.Sprint(hg.TeamId),
					joinUint32(hg.FeaturesId),
					joinUint32(hg.TagsId),
					joinUint32(hg.ShareProductsId),
					joinUint32(hg.ShareTeamsId),
				})
			}
			table.Render()
		case "text":
			if len(allHostgroups) == 0 {
				fmt.Println("No hostgroups found")
				return
			}
			for _, hg := range allHostgroups {
				fmt.Printf("ID: %d, Name: %s, Description: %s, Cluster: %d, DC: %d, Env: %d, Product: %d, Team: %d, Features: [%s], Tags: [%s], ShareProducts: [%s], ShareTeams: [%s]\n",
					hg.Id, hg.Name, hg.Description, hg.ClusterId, hg.DatacenterId, hg.EnvId, hg.ProductId, hg.TeamId,
					joinUint32(hg.FeaturesId),
					joinUint32(hg.TagsId),
					joinUint32(hg.ShareProductsId),
					joinUint32(hg.ShareTeamsId))
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getHostgroupCmd)

	// 添加过滤参数
	getHostgroupCmd.Flags().StringSlice("names", []string{}, "Filter by hostgroup names")
	getHostgroupCmd.Flags().UintSlice("ids", []uint{}, "Filter by hostgroup IDs")
	getHostgroupCmd.Flags().UintSlice("clusters", []uint{}, "Filter by cluster IDs")
	getHostgroupCmd.Flags().UintSlice("datacenters", []uint{}, "Filter by datacenter IDs")
	getHostgroupCmd.Flags().UintSlice("envs", []uint{}, "Filter by environment IDs")
	getHostgroupCmd.Flags().UintSlice("products", []uint{}, "Filter by product IDs")
	getHostgroupCmd.Flags().UintSlice("teams", []uint{}, "Filter by team IDs")
	getHostgroupCmd.Flags().UintSlice("features", []uint{}, "Filter by feature IDs")
	getHostgroupCmd.Flags().UintSlice("tags", []uint{}, "Filter by tag IDs")
	getHostgroupCmd.Flags().UintSlice("share-products", []uint{}, "Filter by shared product IDs")
	getHostgroupCmd.Flags().UintSlice("share-teams", []uint{}, "Filter by shared team IDs")
}

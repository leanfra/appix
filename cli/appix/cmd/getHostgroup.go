package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

		// convert to readableHostgroups
		// Create clients for related services
		clustersClient := pb.NewClustersClient(conn)
		datacentersClient := pb.NewDatacentersClient(conn)
		envsClient := pb.NewEnvsClient(conn)
		productsClient := pb.NewProductsClient(conn)
		teamsClient := pb.NewTeamsClient(conn)
		featuresClient := pb.NewFeaturesClient(conn)
		tagsClient := pb.NewTagsClient(conn)

		// Create caches for related data
		clusterCache := make(map[uint32]string)
		datacenterCache := make(map[uint32]string)
		envCache := make(map[uint32]string)
		productCache := make(map[uint32]string)
		teamCache := make(map[uint32]string)
		featureCache := make(map[uint32]string)
		tagCache := make(map[uint32]string)

		// Collect all unique IDs that need to be looked up
		clusterIDs := make(map[uint32]bool)
		datacenterIDs := make(map[uint32]bool)
		envIDs := make(map[uint32]bool)
		productIDs := make(map[uint32]bool)
		teamIDs := make(map[uint32]bool)
		featureIDs := make(map[uint32]bool)
		tagIDs := make(map[uint32]bool)
		shareProductIDs := make(map[uint32]bool)
		shareTeamIDs := make(map[uint32]bool)

		for _, hg := range allHostgroups {
			if hg.ClusterId > 0 {
				clusterIDs[hg.ClusterId] = true
			}
			if hg.DatacenterId > 0 {
				datacenterIDs[hg.DatacenterId] = true
			}
			if hg.EnvId > 0 {
				envIDs[hg.EnvId] = true
			}
			if hg.ProductId > 0 {
				productIDs[hg.ProductId] = true
			}
			if hg.TeamId > 0 {
				teamIDs[hg.TeamId] = true
			}
			for _, fID := range hg.FeaturesId {
				featureIDs[fID] = true
			}
			for _, tID := range hg.TagsId {
				tagIDs[tID] = true
			}
			for _, pID := range hg.ShareProductsId {
				shareProductIDs[pID] = true
			}
			for _, tID := range hg.ShareTeamsId {
				shareTeamIDs[tID] = true
			}
		}

		ctx := context.Background()
		// Batch fetch clusters
		for id := range clusterIDs {
			resp, err := clustersClient.GetClusters(ctx, &pb.GetClustersRequest{Id: id})
			if err == nil && resp.Cluster != nil {
				clusterCache[id] = resp.Cluster.Name
			} else {
				clusterCache[id] = fmt.Sprint(id)
			}
		}

		// Batch fetch datacenters
		for id := range datacenterIDs {
			resp, err := datacentersClient.GetDatacenters(ctx, &pb.GetDatacentersRequest{Id: id})
			if err == nil && resp.Datacenter != nil {
				datacenterCache[id] = resp.Datacenter.Name
			} else {
				datacenterCache[id] = fmt.Sprint(id)
			}
		}

		// Batch fetch envs
		for id := range envIDs {
			resp, err := envsClient.GetEnvs(ctx, &pb.GetEnvsRequest{Id: id})
			if err == nil && resp.Env != nil {
				envCache[id] = resp.Env.Name
			} else {
				envCache[id] = fmt.Sprint(id)
			}
		}

		// Batch fetch products
		for id := range productIDs {
			resp, err := productsClient.GetProducts(ctx, &pb.GetProductsRequest{Id: id})
			if err == nil && resp.Product != nil {
				productCache[id] = resp.Product.Name
			} else {
				productCache[id] = fmt.Sprint(id)
			}
		}

		// Also fetch share products
		for id := range shareProductIDs {
			if _, exists := productCache[id]; !exists {
				resp, err := productsClient.GetProducts(ctx, &pb.GetProductsRequest{Id: id})
				if err == nil && resp.Product != nil {
					productCache[id] = resp.Product.Name
				} else {
					productCache[id] = fmt.Sprint(id)
				}
			}
		}

		// Batch fetch teams
		for id := range teamIDs {
			resp, err := teamsClient.GetTeams(ctx, &pb.GetTeamsRequest{Id: id})
			if err == nil && resp.Team != nil {
				teamCache[id] = resp.Team.Name
			} else {
				teamCache[id] = fmt.Sprint(id)
			}
		}

		// Also fetch share teams
		for id := range shareTeamIDs {
			if _, exists := teamCache[id]; !exists {
				resp, err := teamsClient.GetTeams(ctx, &pb.GetTeamsRequest{Id: id})
				if err == nil && resp.Team != nil {
					teamCache[id] = resp.Team.Name
				} else {
					teamCache[id] = fmt.Sprint(id)
				}
			}
		}

		// Batch fetch features
		for id := range featureIDs {
			resp, err := featuresClient.GetFeatures(ctx, &pb.GetFeaturesRequest{Id: id})
			if err == nil && resp.Feature != nil {
				featureCache[id] = fmt.Sprintf("%s:%s", resp.Feature.Name, resp.Feature.Value)
			} else {
				featureCache[id] = fmt.Sprint(id)
			}
		}

		// Batch fetch tags
		for id := range tagIDs {
			resp, err := tagsClient.GetTags(ctx, &pb.GetTagsRequest{Id: id})
			if err == nil && resp.Tag != nil {
				tagCache[id] = fmt.Sprintf("%s:%s", resp.Tag.Key, resp.Tag.Value)
			} else {
				tagCache[id] = fmt.Sprint(id)
			}
		}

		var readableHostgroups []*pb.HostgroupReadable
		for _, hg := range allHostgroups {
			readable := &pb.HostgroupReadable{
				Id:            hg.Id,
				Name:          hg.Name,
				Description:   hg.Description,
				Features:      make([]string, len(hg.FeaturesId)),
				Tags:          make([]string, len(hg.TagsId)),
				ShareProducts: make([]string, len(hg.ShareProductsId)),
				ShareTeams:    make([]string, len(hg.ShareTeamsId)),
			}

			readable.Cluster = clusterCache[hg.ClusterId]
			readable.Datacenter = datacenterCache[hg.DatacenterId]
			readable.Env = envCache[hg.EnvId]
			readable.Product = productCache[hg.ProductId]
			readable.Team = teamCache[hg.TeamId]

			for i, id := range hg.FeaturesId {
				readable.Features[i] = featureCache[id]
			}

			for i, id := range hg.TagsId {
				readable.Tags[i] = tagCache[id]
			}

			for i, id := range hg.ShareProductsId {
				readable.ShareProducts[i] = productCache[id]
			}

			for i, id := range hg.ShareTeamsId {
				readable.ShareTeams[i] = teamCache[id]
			}

			readableHostgroups = append(readableHostgroups, readable)
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
			table.SetHeader([]string{"ID", "Name", "Description", "Cluster", "Datacenter", "Env", "Product", "Team", "Features", "Tags", "ShareProducts", "ShareTeams"})
			table.SetAutoFormatHeaders(false)
			for _, hg := range readableHostgroups {
				table.Append([]string{
					fmt.Sprint(hg.Id),
					hg.Name,
					hg.Description,
					hg.Cluster,
					hg.Datacenter,
					hg.Env,
					hg.Product,
					hg.Team,
					strings.Join(hg.Features, ", "),
					strings.Join(hg.Tags, ", "),
					strings.Join(hg.ShareProducts, ", "),
					strings.Join(hg.ShareTeams, ", "),
				})
			}
			table.Render()
		case "text":
			if len(allHostgroups) == 0 {
				fmt.Println("No hostgroups found")
				return
			}
			for _, hg := range readableHostgroups {
				fmt.Printf("ID:            %d\n", hg.Id)
				fmt.Printf("Name:          %s\n", hg.Name)
				fmt.Printf("Description:   %s\n", hg.Description)
				fmt.Printf("Cluster:       %s\n", hg.Cluster)
				fmt.Printf("Datacenter:    %s\n", hg.Datacenter)
				fmt.Printf("Env:           %s\n", hg.Env)
				fmt.Printf("Product:       %s\n", hg.Product)
				fmt.Printf("Team:          %s\n", hg.Team)
				fmt.Printf("Features:      [%s]\n", strings.Join(hg.Features, ", "))
				fmt.Printf("Tags:          [%s]\n", strings.Join(hg.Tags, ", "))
				fmt.Printf("ShareProducts: [%s]\n", strings.Join(hg.ShareProducts, ", "))
				fmt.Printf("ShareTeams:    [%s]\n", strings.Join(hg.ShareTeams, ", "))
				fmt.Println()
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

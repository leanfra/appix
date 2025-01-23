package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "opspillar/api/opspillar/v1"
)

var getAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Get applications",
	Long: `Get applications resources from the system.

Examples:
  opspillar get app                                    # List all
  opspillar get app --names app1,app2                  # Filter by names
  opspillar get app --ids 1,2,3                       # Filter by IDs
  opspillar get app --is-stateful true                # Filter stateful
  opspillar get app --page 1 --page-size 10           # With pagination
  opspillar get app --names web --clusters 1 --format yaml   # Combined filters`,
	Aliases: []string{"apps", "applications"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
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

		// 转换为ApplicationReadable
		// Create clients for related services
		productsClient := pb.NewProductsClient(conn)
		teamsClient := pb.NewTeamsClient(conn)
		hostgroupsClient := pb.NewHostgroupsClient(conn)
		featuresClient := pb.NewFeaturesClient(conn)
		tagsClient := pb.NewTagsClient(conn)
		admClient := pb.NewAdminClient(conn)

		// Create caches for related data
		productCache := make(map[uint32]string)
		teamCache := make(map[uint32]string)
		hostgroupCache := make(map[uint32]string)
		featureCache := make(map[uint32]string)
		tagCache := make(map[uint32]string)
		ownerCache := make(map[uint32]string)

		// Collect all unique IDs that need to be looked up
		productIDs := make(map[uint32]bool)
		teamIDs := make(map[uint32]bool)
		hostgroupIDs := make(map[uint32]bool)
		featureIDs := make(map[uint32]bool)
		tagIDs := make(map[uint32]bool)
		ownerIDs := make(map[uint32]bool)

		for _, app := range allApps {
			if app.ProductId > 0 {
				productIDs[app.ProductId] = true
			}
			if app.TeamId > 0 {
				teamIDs[app.TeamId] = true
			}
			if app.OwnerId > 0 {
				ownerIDs[app.OwnerId] = true
			}
			for _, hgID := range app.HostgroupsId {
				hostgroupIDs[hgID] = true
			}
			for _, fID := range app.FeaturesId {
				featureIDs[fID] = true
			}
			for _, tID := range app.TagsId {
				tagIDs[tID] = true
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

		// Batch fetch teams
		for id := range teamIDs {
			resp, err := teamsClient.GetTeams(ctx, &pb.GetTeamsRequest{Id: id})
			if err == nil && resp.Team != nil {
				teamCache[id] = resp.Team.Name
			} else {
				teamCache[id] = fmt.Sprint(id)
			}
		}

		// Batch fetch owners
		for id := range ownerIDs {
			resp, err := admClient.GetUsers(ctx, &pb.GetUsersRequest{Id: id})
			if err == nil && resp.User != nil {
				ownerCache[id] = resp.User.UserName
			} else {
				ownerCache[id] = fmt.Sprint(id)
			}
		}

		// Batch fetch hostgroups
		for id := range hostgroupIDs {
			resp, err := hostgroupsClient.GetHostgroups(ctx, &pb.GetHostgroupsRequest{Id: id})
			if err == nil && resp.Hostgroup != nil {
				hostgroupCache[id] = resp.Hostgroup.Name
			} else {
				hostgroupCache[id] = fmt.Sprint(id)
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

		var readableApps []*pb.ApplicationReadable
		for _, app := range allApps {
			readable := &pb.ApplicationReadable{
				Id:          app.Id,
				Name:        app.Name,
				Description: app.Description,
				IsStateful:  app.IsStateful,
				Features:    make([]string, len(app.FeaturesId)),
				Tags:        make([]string, len(app.TagsId)),
				Hostgroups:  make([]string, len(app.HostgroupsId)),
			}

			// Use cached product name
			readable.Product = productCache[app.ProductId]

			// Use cached team name
			readable.Team = teamCache[app.TeamId]

			// Use cached owner name
			readable.Owner = ownerCache[app.OwnerId]

			// Use cached hostgroup names
			for i, hgID := range app.HostgroupsId {
				readable.Hostgroups[i] = hostgroupCache[hgID]
			}

			// Convert features to "name:value" format using cached names
			for i, id := range app.FeaturesId {
				readable.Features[i] = featureCache[id]
			}

			// Convert tags to "name:value" format using cached names
			for i, id := range app.TagsId {
				readable.Tags[i] = tagCache[id]
			}

			readableApps = append(readableApps, readable)
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
			table.SetAutoFormatHeaders(false)
			for _, app := range readableApps {
				table.Append([]string{
					fmt.Sprint(app.Id),
					app.Name,
					app.Description,
					app.Owner,
					fmt.Sprint(app.IsStateful),
					app.Product,
					app.Team,
					strings.Join(app.Features, ", "),
					strings.Join(app.Tags, ", "),
					strings.Join(app.Hostgroups, ", "),
				})
			}
			table.Render()

		case "text":
			if len(readableApps) == 0 {
				fmt.Println("No applications found")
				return
			}
			for _, app := range readableApps {
				fmt.Printf("ID:          %d\n"+
					"Name:        %s\n"+
					"Description: %s\n"+
					"Owner:       %s\n"+
					"Stateful:    %v\n"+
					"Product:     %s\n"+
					"Team:        %s\n"+
					"Features:    %s\n"+
					"Tags:        %s\n"+
					"Hostgroups:  %s\n\n",
					app.Id, app.Name, app.Description, app.Owner,
					app.IsStateful, app.Product, app.Team,
					strings.Join(app.Features, ", "),
					strings.Join(app.Tags, ", "),
					strings.Join(app.Hostgroups, ", "))
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

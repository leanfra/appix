/*
Copyright © 2024 Lenfra <lenfra@163.com>
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

// getTeamCmd represents the getTeam command
var getTeamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get teams",
	Long: `Get teams resources from the system.

This command allows you to retrieve and list teams with various filtering options:
  - Filter by team names (-N, --names)
  - Filter by team codes (-C, --codes)
  - Filter by team leaders (-L, --leaders)
  - Filter by team IDs (-I, --ids)

Pagination support:
  - page: Specify the page number
  - page-size: Number of items to display per page

Output formats available:
  - table: Displays results in a formatted table
  - yaml: Outputs in YAML format
  - text: Simple text output format

Examples:
  appix get team                      # List all teams
  appix get team -N team1,team2       # Filter teams by names
  appix get team -C dev,prod          # Filter teams by codes
  appix get team -L alice,bob         # Filter teams by leaders
  appix get team -I 1,2,3             # Filter teams by IDs
  appix get team --page 1 --page-size 10  # Paginated results`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("连接失败: %v", err)
		}
		defer conn.Close()

		client := pb.NewTeamsClient(conn)

		uint32Ids := make([]uint32, len(teamIds))
		for i, id := range teamIds {
			uint32Ids[i] = uint32(id)
		}

		filter := &pb.ListTeamsFilter{
			Page:     GetPage,
			PageSize: GetPageSize,
			Names:    teamNames,
			Codes:    teamCodes,
			Leaders:  teamLeaders,
			Ids:      uint32Ids,
		}

		req := &pb.ListTeamsRequest{
			Filter: filter,
		}

		ctx := context.Background()
		reply, err := client.ListTeams(ctx, req)
		if err != nil {
			fmt.Printf("Failed to get teams: %v\n", err)
			return
		}

		if reply.Code != 0 {
			fmt.Printf("Error: %s\n", reply.Message)
			return
		}

		// 按 GetFormat 格式输出
		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(reply.Teams)
			if err != nil {
				log.Fatalf("序列化YAML失败: %v", err)
			}
			fmt.Println(string(data))
		case "text":
			if len(reply.Teams) == 0 {
				fmt.Println("No teams found")
				return
			}
			for _, team := range reply.Teams {
				fmt.Printf("ID: %d, Name: %s, Code: %s, Leader: %s, Description: %s\n", team.Id, team.Name, team.Code, team.Leader, team.Description)
			}
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Code", "Leader", "Description"})
			for _, team := range reply.Teams {
				table.Append([]string{
					fmt.Sprint(team.Id),
					team.Name,
					team.Code,
					team.Leader,
					team.Description,
				})
			}
			table.Render()
		default:
			fmt.Println("unknown format")
		}
	},
}

var teamNames []string
var teamCodes []string
var teamLeaders []string
var teamIds []uint

func init() {
	getCmd.AddCommand(getTeamCmd)

	getTeamCmd.Flags().StringSliceVarP(&teamNames, "names", "N", []string{}, "Filter by team names (comma-separated)")
	getTeamCmd.Flags().StringSliceVarP(&teamCodes, "codes", "C", []string{}, "Filter by team codes (comma-separated)")
	getTeamCmd.Flags().StringSliceVarP(&teamLeaders, "leaders", "L", []string{}, "Filter by team leaders (comma-separated)")
	getTeamCmd.Flags().UintSliceVarP(&teamIds, "ids", "I", []uint{}, "Filter by team IDs (comma-separated)")
}

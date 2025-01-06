/*
Copyright © 2024 Lenfra <lenfra@163.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// getTeamCmd represents the getTeam command
var getTeamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get teams",
	Long: `Get teams resources from the system.

Examples:
  appix get team                              # List all
  appix get team --names team1,team2          # Filter by names
  appix get team --codes dev,prod             # Filter by codes
  appix get team --leaders alice,bob          # Filter by leaders
  appix get team --names dev --format yaml    # Custom format`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewTeamsClient(conn)

		uint32Ids := make([]uint32, len(teamIds))
		for i, id := range teamIds {
			uint32Ids[i] = uint32(id)
		}

		// 创建一个切片存储所有团队
		var allTeams []*pb.Team
		currentPage := GetPage

		for {

			req := &pb.ListTeamsRequest{
				Page:     currentPage,
				PageSize: GetPageSize,
				Names:    teamNames,
				Codes:    teamCodes,
				Leaders:  teamLeaders,
				Ids:      uint32Ids,
			}

			reply, err := client.ListTeams(ctx, req)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if reply.Code != 0 {
				fmt.Printf("Response details:\n")
				fmt.Printf("  Message: %s\n", reply.Message)
				fmt.Printf("  Code: %d\n", reply.Code)
				fmt.Printf("  Action: %s\n", reply.Action)
				return
			}

			// 添加当前页的团队到总列表
			allTeams = append(allTeams, reply.Teams...)

			// 如果返回的团队数量小于页大小，说明已经是最后一页
			if len(reply.Teams) < int(GetPageSize) {
				break
			}

			currentPage++
		}

		// 按 GetFormat 格式输出
		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allTeams)
			if err != nil {
				log.Fatalf("serialize yaml failed: %v", err)
			}
			fmt.Println(string(data))
		case "text":
			if len(allTeams) == 0 {
				fmt.Println("No teams found")
				return
			}
			for _, team := range allTeams {
				fmt.Printf("ID: %d, Name: %s, Code: %s, Leader: %s, Description: %s\n",
					team.Id, team.Name, team.Code, team.Leader, team.Description)
			}
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Code", "Leader", "Description"})
			for _, team := range allTeams {
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

	getTeamCmd.Flags().StringSlice("names", []string{}, "Filter by team names (comma-separated)")

	getTeamCmd.Flags().StringSlice("codes", []string{}, "Filter by team codes (comma-separated)")

	getTeamCmd.Flags().StringSlice("leaders", []string{}, "Filter by team leaders (comma-separated)")

	getTeamCmd.Flags().UintSlice("ids", []uint{}, "Filter by team IDs (comma-separated)")
}

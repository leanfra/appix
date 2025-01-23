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

	pb "opspillar/api/opspillar/v1"
)

// getTeamCmd represents the getTeam command
var getTeamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get teams",
	Long: `Get teams resources from the system.

Examples:
  opspillar get team                              # List all
  opspillar get team --names team1,team2          # Filter by names
  opspillar get team --codes dev,prod             # Filter by codes
  opspillar get team --leaders 1,2 				  # Filter by leaders user id
  opspillar get team --names dev --format yaml    # Custom format`,
	Aliases: []string{"teams", "tm"},
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
				Page:      currentPage,
				PageSize:  GetPageSize,
				Names:     teamNames,
				Codes:     teamCodes,
				LeadersId: teamLeadersId,
				Ids:       uint32Ids,
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

		// convert teams to TeamReadable
		// 创建一个缓存来存储 leader_id 对应的 username
		leaderCache := make(map[uint32]string)

		// 收集所有唯一的 leader_id 需要被查询
		leaderIDs := make(map[uint32]bool)
		for _, team := range allTeams {
			if team.LeaderId > 0 {
				leaderIDs[team.LeaderId] = true
			}
		}

		// 批量获取 leaders
		adminClient := pb.NewAdminClient(conn)

		for id := range leaderIDs {
			resp, err := adminClient.GetUsers(ctx, &pb.GetUsersRequest{Id: id})
			if err == nil && resp.User != nil {
				leaderCache[id] = resp.User.UserName
			} else {
				leaderCache[id] = fmt.Sprint(id)
			}
		}

		// 将 pb.Team 转换为 pb.TeamReadable
		var readableTeams []*pb.TeamReadable
		for _, team := range allTeams {
			readable := &pb.TeamReadable{
				Id:          team.Id,
				Name:        team.Name,
				Code:        team.Code,
				Description: team.Description,
				Leader:      leaderCache[team.LeaderId],
			}
			readableTeams = append(readableTeams, readable)
		}

		// 按 GetFormat 格式输出
		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(readableTeams)
			if err != nil {
				log.Fatalf("serialize yaml failed: %v", err)
			}
			fmt.Println(string(data))
		case "text":
			if len(readableTeams) == 0 {
				fmt.Println("No teams found")
				return
			}
			for _, team := range readableTeams {
				fmt.Printf("ID: %d, Name: %s, Code: %s, Leader: %s, Description: %s\n",
					team.Id, team.Name, team.Code, team.Leader, team.Description)
			}
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Code", "Leader", "Description"})
			table.SetAutoFormatHeaders(true)
			for _, team := range readableTeams {
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
var teamLeadersId []uint32
var teamIds []uint

func init() {
	getCmd.AddCommand(getTeamCmd)

	getTeamCmd.Flags().StringSlice("names", []string{}, "Filter by team names (comma-separated)")

	getTeamCmd.Flags().StringSlice("codes", []string{}, "Filter by team codes (comma-separated)")

	getTeamCmd.Flags().UintSlice("leaders-id", []uint{}, "Filter by team leaders (comma-separated)")

	getTeamCmd.Flags().UintSlice("ids", []uint{}, "Filter by team IDs (comma-separated)")
}

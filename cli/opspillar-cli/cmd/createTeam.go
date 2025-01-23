/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	pb "opspillar/api/opspillar/v1"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// createTeamCmd represents the createTeam command
var createTeamCmd = &cobra.Command{
	Use:   "team",
	Short: "Create a new team",
	Long: `Create a new team in the system.
Team is a development team as organization unit.

Examples:
  opspillar create team --name eng-team --code eng --desc "Engineering Team" --leader "John Doe"
  opspillar create team --name ops-team --code ops --desc "Operations Team" --leader "Jane Smith"`,
	Aliases: []string{"tm", "teams"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewTeamsClient(conn)

		var req *pb.CreateTeamsRequest

		if outFile != "" {
			// Generate template YAML file
			team := &pb.Team{
				Name:        "team-name",
				Code:        "team-code",
				Description: "team description",
				LeaderId:    1,
			}
			teams := []*pb.Team{team}

			data, err := yaml.Marshal(teams)
			if err != nil {
				log.Fatalf("failed to generate yaml: %v", err)
			}

			if err := os.WriteFile(outFile, data, 0644); err != nil {
				log.Fatalf("failed to write template file: %v", err)
			}

			fmt.Printf("Template file generated at: %s\n", outFile)
			return
		} else if yamlFile != "" {
			// Read from YAML file
			data, err := os.ReadFile(yamlFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			var teams []*pb.Team
			if err := yaml.Unmarshal(data, &teams); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateTeamsRequest{
				Teams: teams,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			code, _ := cmd.Flags().GetString("code")
			desc, _ := cmd.Flags().GetString("desc")
			leaderId, _ := cmd.Flags().GetUint32("leader")

			req = &pb.CreateTeamsRequest{
				Teams: []*pb.Team{
					{
						Name:        name,
						Code:        code,
						Description: desc,
						LeaderId:    leaderId,
					},
				},
			}
		}

		resp, err := client.CreateTeams(ctx, req)
		if err != nil {
			log.Fatalf("failed to create teams: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}

	},
}

func init() {
	createCmd.AddCommand(createTeamCmd)
	createTeamCmd.Flags().String("name", "", "Name of the team")
	createTeamCmd.Flags().String("code", "", "Code of the team")
	createTeamCmd.Flags().String("desc", "", "Description of the team")
	createTeamCmd.Flags().Uint32("leaderId", 0, "UserId of the Leader of the team")
}

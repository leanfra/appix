/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "appix/api/appix/v1"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// updateTeamCmd represents the updateTeam command
var updateTeamCmd = &cobra.Command{
	Use:     "team",
	Short:   "Update a team",
	Long:    `Update a team with the specified ID and fields.`,
	Aliases: []string{"team", "teams", "tm"},
	Run: func(cmd *cobra.Command, args []string) {

		ctx, conn, err := NewConnection(true)
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			return
		}
		defer conn.Close()

		client := pb.NewTeamsClient(conn)

		yamlFile, _ := cmd.Flags().GetString("yaml")
		editOnline, _ := cmd.Flags().GetBool("edit")
		var teams []*pb.Team
		if editOnline {
			// Get existing team data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the team data
			getReq := &pb.GetTeamsRequest{Id: id}
			getResp, err := client.GetTeams(ctx, getReq)
			if err != nil {
				log.Fatalf("failed to get team: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Team{getResp.Team})
			if err != nil {
				log.Fatalf("failed to marshal team: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "team-*.yaml")
			if err != nil {
				log.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write(data); err != nil {
				log.Fatalf("failed to write temp file: %v", err)
			}
			tmpfile.Close()

			editor := findEditor()
			if editor == "" {
				log.Fatal("no suitable editor found - please set EDITOR environment variable")
			}

			cmd := exec.Command(editor, tmpfile.Name())
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("failed to run editor: %v", err)
			}

			// Read updated content
			updatedData, err := os.ReadFile(tmpfile.Name())
			if err != nil {
				log.Fatalf("failed to read updated file: %v", err)
			}
			if string(updatedData) == string(data) {
				fmt.Println("No changes detected, skipping update")
				return
			}

			// Parse updated YAML
			if err := yaml.Unmarshal(updatedData, &teams); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if yamlFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(yamlFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &teams); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			name, _ := cmd.Flags().GetString("name")
			description, _ := cmd.Flags().GetString("description")
			code, _ := cmd.Flags().GetString("code")
			leaderId, _ := cmd.Flags().GetUint32("leader")

			teams = []*pb.Team{
				{
					Id:          id,
					Name:        name,
					Description: description,
					Code:        code,
					LeaderId:    leaderId,
				},
			}
		}

		req := &pb.UpdateTeamsRequest{
			Teams: teams,
		}

		reply, err := client.UpdateTeams(ctx, req)
		if err != nil {
			log.Fatalf("failed to update team: %v", err)
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateTeamCmd)

	updateTeamCmd.Flags().Uint32("id", 0, "Team ID to update")
	updateTeamCmd.Flags().String("code", "", "New team code")
	updateTeamCmd.Flags().Uint32("leader-id", 0, "New team leader user id")
	updateTeamCmd.Flags().String("name", "", "New team name")
	updateTeamCmd.Flags().String("desc", "", "New team description")

}

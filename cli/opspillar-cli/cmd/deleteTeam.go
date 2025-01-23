/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	pb "opspillar/api/opspillar/v1"
)

// deleteTeamCmd represents the deleteTeam command
var deleteTeamCmd = &cobra.Command{
	Use:   "team [ids...]",
	Short: "Delete one or more teams by their IDs",
	Long: `Delete one or more teams by providing their IDs as arguments.
For example:
  opspillar delete team 1 2 3`,
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"team", "teams", "tm"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewTeamsClient(conn)

		if len(args) == 0 {
			fmt.Println("Please provide at least one team ID")
			return
		}

		ids := make([]uint32, 0, len(args))
		for _, arg := range args {
			var id uint64
			id, err := strconv.ParseUint(arg, 10, 32)
			if err != nil {
				fmt.Printf("Invalid team ID '%s': %v\n", arg, err)
				return
			}
			ids = append(ids, uint32(id))
		}

		req := &pb.DeleteTeamsRequest{
			Ids: ids,
		}

		reply, err := client.DeleteTeams(ctx, req)
		if err != nil {
			log.Fatalf("failed to delete teams: %v", err)
		}

		// print reply code, message, and so on
		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}

	},
}

func init() {
	deleteCmd.AddCommand(deleteTeamCmd)
}

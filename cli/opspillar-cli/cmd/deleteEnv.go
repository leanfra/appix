/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	pb "opspillar/api/opspillar/v1"

	"github.com/spf13/cobra"
)

// deleteEnvCmd represents the deleteEnv command
var deleteEnvCmd = &cobra.Command{
	Use:   "env [ids...]",
	Short: "Delete one or more environments by their IDs",
	Long: `Delete one or more environments by providing their IDs as arguments.
For example:
  opspillar delete env 1 2 3`,
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"env", "envs", "en"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewEnvsClient(conn)

		if len(args) == 0 {
			fmt.Println("Please provide at least one environment ID")
			return
		}

		ids := make([]uint32, 0, len(args))
		for _, arg := range args {
			id, err := strconv.ParseUint(arg, 10, 32)
			if err != nil {
				fmt.Printf("Invalid environment ID '%s': %v\n", arg, err)
				return
			}
			ids = append(ids, uint32(id))
		}

		req := &pb.DeleteEnvsRequest{
			Ids: ids,
		}

		reply, err := client.DeleteEnvs(ctx, req)
		if err != nil {
			log.Fatalf("failed to delete environments: %v", err)
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteEnvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteEnvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

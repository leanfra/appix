/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	pb "appix/api/appix/v1"

	"github.com/spf13/cobra"
)

// deleteDatacenterCmd represents the deleteDatacenter command
var deleteDatacenterCmd = &cobra.Command{
	Use:   "datacenter [ids...]",
	Short: "Delete one or more datacenters by their IDs",
	Long: `Delete one or more datacenters by providing their IDs as arguments.
For example:
  appix delete datacenter 1 2 3`,
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"datacenters", "dc"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewDatacentersClient(conn)

		if len(args) == 0 {
			fmt.Println("Please provide at least one datacenter ID")
			return
		}

		ids := make([]uint32, 0, len(args))
		for _, arg := range args {
			id, err := strconv.ParseUint(arg, 10, 32)
			if err != nil {
				fmt.Printf("Invalid datacenter ID '%s': %v\n", arg, err)
				return
			}
			ids = append(ids, uint32(id))
		}

		req := &pb.DeleteDatacentersRequest{
			Ids: ids,
		}

		reply, err := client.DeleteDatacenters(ctx, req)
		if err != nil {
			log.Fatalf("failed to delete datacenters: %v", err)
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteDatacenterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDatacenterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDatacenterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

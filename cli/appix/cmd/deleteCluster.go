/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	pb "appix/api/appix/v1"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// deleteClusterCmd represents the deleteCluster command
var deleteClusterCmd = &cobra.Command{
	Use:   "cluster [ids...]",
	Short: "Delete one or more clusters by their IDs",
	Long: `Delete one or more clusters by providing their IDs as arguments.
For example:
  appix delete cluster 1 2 3`,
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"cluster", "clusters", "cl"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewClustersClient(conn)

		if len(args) == 0 {
			fmt.Println("Please provide at least one cluster ID")
			return
		}

		ids := make([]uint32, 0, len(args))
		for _, arg := range args {
			id, err := strconv.ParseUint(arg, 10, 32)
			if err != nil {
				fmt.Printf("Invalid cluster ID '%s': %v\n", arg, err)
				return
			}
			ids = append(ids, uint32(id))
		}

		req := &pb.DeleteClustersRequest{
			Ids: ids,
		}

		reply, err := client.DeleteClusters(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error deleting clusters: %v\n", err)
			return
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteClusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteClusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

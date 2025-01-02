/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "appix/api/appix/v1"
)

// deleteFeatureCmd represents the deleteFeature command
var deleteFeatureCmd = &cobra.Command{
	Use:   "feature [ids...]",
	Short: "Delete one or more features by their IDs",
	Long: `Delete one or more features by providing their IDs as arguments.
For example:
  appix delete feature 1 2 3`,
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"feature", "features", "ft"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewFeaturesClient(conn)

		if len(args) == 0 {
			fmt.Println("Please provide at least one feature ID")
			return
		}

		ids := make([]uint32, 0, len(args))
		for _, arg := range args {
			id, err := strconv.ParseUint(arg, 10, 32)
			if err != nil {
				fmt.Printf("Invalid feature ID '%s': %v\n", arg, err)
				return
			}
			ids = append(ids, uint32(id))
		}

		req := &pb.DeleteFeaturesRequest{
			Ids: ids,
		}

		reply, err := client.DeleteFeatures(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error deleting features: %v\n", err)
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
	deleteCmd.AddCommand(deleteFeatureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteFeatureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteFeatureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

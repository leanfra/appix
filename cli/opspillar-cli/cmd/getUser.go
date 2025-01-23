/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	v1 "opspillar/api/opspillar/v1"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// getUserCmd represents the getUser command
var getUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Get users",
	Long: `Get users from the system.

Examples:
  opspillar get user                                    # List all
  opspillar get user --names user1,user2                  # Filter by names
  opspillar get user --ids 1,2,3                       # Filter by IDs
  opspillar get user --page 1 --page-size 10           # With pagination
  opspillar get user --names john --format yaml          # Combined filters`,
	Aliases: []string{"users", "usr"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := v1.NewAdminClient(conn)

		page := GetPage
		pageSize := GetPageSize
		names, _ := cmd.Flags().GetStringSlice("names")

		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		ids := toUint32Slice(uintIds)

		phones, _ := cmd.Flags().GetStringSlice("phone")
		emails, _ := cmd.Flags().GetStringSlice("email")

		var allUsers []*v1.User
		for {
			req := &v1.ListUsersRequest{
				Page:      page,
				PageSize:  pageSize,
				UserNames: names,
				Ids:       ids,
				Phones:    phones,
				Emails:    emails,
			}

			resp, err := client.ListUsers(ctx, req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			if resp.Code != 0 {
				fmt.Printf("Response details:\n")
				fmt.Printf("  Message: %s\n", resp.Message)
				fmt.Printf("  Code: %d\n", resp.Code)
				fmt.Printf("  Action: %s\n", resp.Action)
				return
			}
			allUsers = append(allUsers, resp.Users...)

			if len(resp.Users) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allUsers)
			if err != nil {
				log.Fatalf("failed to generate yaml: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{
				"ID", "Name", "Email", "Phone", "Password",
			})
			table.SetAutoFormatHeaders(true)
			for _, user := range allUsers {
				table.Append([]string{
					fmt.Sprint(user.Id),
					user.UserName,
					user.Email,
					user.Phone,
					user.Password,
				})
			}
			table.Render()

		case "text":
			if len(allUsers) == 0 {
				fmt.Println("No users found")
				return
			}
			for _, user := range allUsers {
				fmt.Printf("ID:          %d\n"+
					"Name:        %s\n"+
					"Email:       %s\n"+
					"Phone:       %s\n\n",
					user.Id, user.UserName, user.Email, user.Phone)
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getUserCmd)

	// 只使用长格式标志
	getUserCmd.Flags().StringSlice("names", []string{}, "Filter by user names")
	getUserCmd.Flags().UintSlice("ids", []uint{}, "Filter by user IDs")
	getUserCmd.Flags().StringSlice("phone", []string{}, "Filter by phone")
	getUserCmd.Flags().StringSlice("email", []string{}, "Filter by email")

}

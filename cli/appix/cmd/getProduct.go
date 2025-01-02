/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "appix/api/appix/v1"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
)

// getProductCmd represents the getProduct command
var getProductCmd = &cobra.Command{
	Use:   "product",
	Short: "Get products resources from the system",
	Long: `Get products resources from the system.

Examples:
  appix get product                              # List all
  appix get product --names prod1,prod2          # Filter by names
  appix get product --codes code1,code2          # Filter by codes
  appix get product --ids 1,2,3                  # Filter by IDs
  appix get product --names prod1 --format yaml  # Custom format`,
	Run: func(cmd *cobra.Command, args []string) {
		names, _ := cmd.Flags().GetStringSlice("names")
		codes, _ := cmd.Flags().GetStringSlice("codes")
		ids, _ := cmd.Flags().GetUintSlice("ids")
		page := GetPage
		pageSize := GetPageSize

		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewProductsClient(conn)

		// 转换 ids 到 uint32
		idsUint32 := make([]uint32, len(ids))
		for i, id := range ids {
			idsUint32[i] = uint32(id)
		}

		// 创建一个切片存储所有产品
		var allProducts []*pb.Product
		currentPage := page

		for {
			req := &pb.ListProductsRequest{
				Filter: &pb.ListProductsFilter{
					Page:     currentPage,
					PageSize: pageSize,
					Names:    names,
					Codes:    codes,
					Ids:      idsUint32,
				},
			}

			ctx := context.Background()
			reply, err := client.ListProducts(ctx, req)

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

			// 添加当前页的产品到总列表
			allProducts = append(allProducts, reply.Products...)

			// 如果返回的产品数量小于页大小，说明已经是最后一页
			if len(reply.Products) < int(pageSize) {
				break
			}

			currentPage++
		}

		// 根据格式输出结果
		switch GetFormat {
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Code", "Description"})

			for _, p := range allProducts {
				table.Append([]string{
					fmt.Sprint(p.Id),
					p.Name,
					p.Code,
					p.Description,
				})
			}
			table.Render()

		case "yaml":
			data, err := yaml.Marshal(allProducts)
			if err != nil {
				log.Fatalf("serialize yaml failed: %v", err)
			}
			fmt.Println(string(data))

		case "text":
			if len(allProducts) == 0 {
				fmt.Println("No products found")
				return
			}
			for _, p := range allProducts {
				fmt.Printf("ID: %d, Name: %s, Code: %s, Description: %s\n",
					p.Id, p.Name, p.Code, p.Description)
			}

		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getProductCmd)

	// Add filter flags
	getProductCmd.Flags().StringSlice("names", []string{}, "Filter by product names")

	getProductCmd.Flags().StringSlice("codes", []string{}, "Filter by product codes")

	getProductCmd.Flags().UintSlice("ids", []uint{}, "Filter by product IDs")

}

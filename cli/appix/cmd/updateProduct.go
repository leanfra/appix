/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// updateProductCmd represents the updateProduct command
var updateProductCmd = &cobra.Command{
	Use:     "product",
	Short:   "Update product",
	Long:    `Update one or more products with the specified ID and fields.`,
	Aliases: []string{"product", "products"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewProductsClient(conn)

		yamlFile, _ := cmd.Flags().GetString("yaml")
		var products []*pb.Product
		if updateOnline {
			// Get existing product data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the product data
			getReq := &pb.GetProductsRequest{Id: id}
			getResp, err := client.GetProducts(cmd.Context(), getReq)
			if err != nil {
				log.Fatalf("failed to get product: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Product{getResp.Product})
			if err != nil {
				log.Fatalf("failed to marshal product: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "product-*.yaml")
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
			if err := yaml.Unmarshal(updatedData, &products); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}
		} else if yamlFile != "" {
			data, err := os.ReadFile(yamlFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &products); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			name, _ := cmd.Flags().GetString("name")
			desc, _ := cmd.Flags().GetString("desc")
			code, _ := cmd.Flags().GetString("code")

			products = []*pb.Product{
				{
					Id:          id,
					Name:        name,
					Description: desc,
					Code:        code,
				},
			}
		}

		req := &pb.UpdateProductsRequest{
			Products: products,
		}

		reply, err := client.UpdateProducts(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error updating product: %v\n", err)
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
	updateCmd.AddCommand(updateProductCmd)

	updateProductCmd.Flags().Uint32("id", 0, "Product ID to update")
	updateProductCmd.Flags().String("name", "", "New product name")
	updateProductCmd.Flags().String("desc", "", "New product description")
	updateProductCmd.Flags().String("code", "", "New product code")
}

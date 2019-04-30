// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"

	blogpb "../../proto"
	"github.com/spf13/cobra"
)

// updateCmd represents the read command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Find a Blog post by its ID",
	Long: `Find a blog post by it's mongoDB Unique identifier.
	
	If no blog post is found for the ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
		author, err := cmd.Flags().GetString("author")
		title, err := cmd.Flags().GetString("title")
		content, err := cmd.Flags().GetString("content")

		// Create an UpdateBlogRequest
		req := &blogpb.UpdateBlogReq{
			&blogpb.Blog{
				Id:       id,
				AuthorId: author,
				Title:    title,
				Content:  content,
			},
		}

		res, err := client.UpdateBlog(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "The id of the blog")
	updateCmd.Flags().StringP("author", "a", "", "Add an author")
	updateCmd.Flags().StringP("title", "t", "", "A title for the blog")
	updateCmd.Flags().StringP("content", "c", "", "The content for the blog")
	updateCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

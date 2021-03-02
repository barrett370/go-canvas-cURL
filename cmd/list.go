/*
Copyright © 2021 Sam Barrett <barrett370@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	lib "github.com/barrett370/go-canvas-cUrl/lib"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all enrolled modules",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		requester, err := lib.GetRequester()
		if err != nil {
			panic(fmt.Errorf("Error getting requester: %s", err))
		}
		courses, err := lib.GetCourses(requester, make([]string, 0))
		if err != nil {
			fmt.Println("ERROR")

		}
		for _, course := range courses {
			fmt.Println("", course.Name)

		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//

}

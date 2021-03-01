/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	lib "github.com/barrett370/go-canvas-cUrl/lib"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
		if len(args) > 0 {
			if strings.ToLower(args[0]) == "all" {
				args = make([]string, 0)
			}
		}

		requester, err := lib.GetRequester()
		if err != nil {
			panic(fmt.Errorf("Error getting requester: %s", err))
		}

		courses, err := lib.GetCourses(requester, args)
		if err != nil {
			panic(fmt.Errorf("Error getting courses %s", err))
		}

		requester.Context = "/api/v1/courses/"
		for _, course := range courses {
			fmt.Println("Searching course: ", course.Name)

			modules, err := course.GetModules(requester)
			if err != nil {
				switch e := err.(type) {
				case *lib.NoModulesError:
					err = course.GetFiles(requester)
					if err != nil {
						fmt.Printf(e.Error() + "\n")
						continue
					}
				case *lib.NoFilesError:
					continue
				default:
					fmt.Printf(e.Error() + "\n")
					continue
				}
			}
			for _, module := range modules {
				folders, err := module.GetFolders(requester)
				if err != nil {
					fmt.Printf(err.Error() + "\n")
				}
				for _, folder := range folders {
					err = folder.GetFiles(requester, course)
					if err != nil {
						fmt.Printf(err.Error() + "\n")
					}
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of past daily reports.",
	Long:  `Write a list of past daily reports to standard output.`,
	Run:   List,
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
}

func List(cmd *cobra.Command, args []string) {
	f, err := os.Open(config.ReportDir)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	files, err := f.Readdirnames(-1)
	if err != nil {
		os.Exit(1)
	}
	files = sortByDate(files)

	for _, file := range files {
		fmt.Println(filepath.Join(config.ReportDir, file))
	}
}

func sortByDate(files []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	return files
}

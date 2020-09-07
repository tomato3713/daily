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
	"github.com/spf13/cobra"
	"os"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Write today's daily report.",
	Long: `This is a command to write a daily report. 
	If the today's daily report has not been created yet, 
	it will create an empty daily report with the contents 
	specified in the template inserted. 
	If you already have a daily report, open it.`,
	Run: report,
}

func report(cmd *cobra.Command, args []string) {
	// make daily report directory if not exists.
	if err := makeDailyReportDirecotry(config.ReportDir); err != nil {
		os.Exit(1)
	}

	// load daily report format file
}

func makeDailyReportDirecotry(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			os.Exit(1)
		}
		err := os.MkdirAll(dir, 0755)
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

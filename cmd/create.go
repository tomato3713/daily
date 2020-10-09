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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an configuration file",
	Long: `Create a configuration file ($HOME/.config/daily/config.toml). 
	It will fail if the config file already exists.`,
	Run: CreateCfgFile,
}

func init() {
	configCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CreateCfgFile(cmd *cobra.Command, args []string) {
	// create configuration file and directory if it is not exists.
	// Find home directory.
	fmt.Println("Create configuration file: ", cfgFile)

	root := filepath.Dir(cfgFile)
	if _, err := os.Stat(root); err != nil {
		if !os.IsNotExist(err) {
			os.Exit(1)
		}
		err := os.MkdirAll(root, 0755)
		if err != nil {
			os.Exit(1)
		}
	}

	_, err := os.Stat(cfgFile)
	confExists := err == nil
	if !confExists {
		// TODO: rewrite
		contents := []byte(fmt.Sprintln(`ReportDir = "$HOME/.config/reports"
FileNameFormat = "${yyyy}-${mm}-${dd}-daily-report.md"
TemplateFile = "$HOME/.config/template.md"
PluginDir = "$HOME/.config/plugins"

[Serve]
TemplateIndexFile = "$HOME/.config/index.tmpl"
TemplateBodyFile = "$HOME/.config/body.tmpl"
AssetsDir = "$HOME/.config/assets"`))

		if err = ioutil.WriteFile(cfgFile, contents, 0644); err != nil {
			os.Exit(1)
		}
	}
	fmt.Println(config)
}

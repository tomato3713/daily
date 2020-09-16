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
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "print daily report you want.",
	Long:  `print daily report you want. Search 'ReportDir' in the configuration file. `,
	Run:   Cat,
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Cat(cmd *cobra.Command, args []string) {
	if len(args) == 3 {
		if err := printFile(args[0], args[1], args[2]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if len(args) != 0 {
		os.Exit(1)
	}

	files, err := selectFile()
	if err != nil {
		os.Exit(1)
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			os.Exit(1)
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(string(b))
	}
}

func selectFile() ([]string, error) {
	var files []string
	f, err := os.Open(config.ReportDir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	files, err = f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = runCmd("fzf", strings.NewReader(strings.Join(files, "\n")), &buf)
	if err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		return nil, errors.New("No selected files")
	}
	selectFiles := strings.Split(strings.TrimSpace(buf.String()), "\n")
	for i, file := range files {
		selectFiles[i] = filepath.Join(config.ReportDir, file)
	}
	return selectFiles, nil
}

func printFile(year, month, day string) error {
	fname := fmt.Sprintf("%s-%s-%s-daily-report.md", year, month, day)
	file := filepath.Join(config.ReportDir, fname)

	if fileExists(file) {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		fmt.Println(string(b))
		return nil
	}
	return fmt.Errorf("Error: Can's open %s", file)
}

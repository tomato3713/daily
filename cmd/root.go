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
	"github.com/spf13/cobra"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Configuration
type Config struct {
	ReportDir      string      // path to daily report file directory
	FileNameFormat string      // daily-report
	TemplateFile   string      // path to template file
	PluginDir      string      // path to plugin directory
	Serve          ServeConfig // for http server
}

type ServeConfig struct {
	TemplateBodyFile string // path to index.html
	AssetsDir        string // path to assets directory
}

var cfgFile string
var debug bool
var config Config

func (c Config) String() string {
	return fmt.Sprintf("Daily Report Dir: %s\nFile Name Format: %s\nDaily report Template File: %s\nPlugins Dir: %s\nServe:\n\tTemplate Body File: %s\n\tAssets Dir: %s\n",
		c.ReportDir, c.FileNameFormat, c.TemplateFile, c.PluginDir, c.Serve.TemplateBodyFile, c.Serve.AssetsDir)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "daily",
	Short: "daily is an application for daily",
	Long: `daily is a command to help manage and write daily reports easily. 
	Daily reports can be written in a markup language like Markdown.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/daily/config.toml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug message")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfgDir := filepath.Join(home, `.config`, `daily`)
		cfgFile = filepath.Join(cfgDir, `config.toml`)

		// Search config in home directory with name ".config/daily/" (without extension).
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

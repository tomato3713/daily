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
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/browser"
	"github.com/shurcooL/github_flavored_markdown"
	"github.com/spf13/cobra"
)

type entry struct {
	Name string
	Body template.HTML
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start http server",
	Long:  `Start http server to display past daily reports.`,
	Run:   Serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Serve(cmd *cobra.Command, args []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			f, err := os.Open(config.ReportDir)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer f.Close()
			files, err := f.Readdirnames(-1)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			files = sortByDate(files)
			var entries []entry
			for _, file := range files {
				entries = append(entries, entry{
					Name: file,
				})
			}
			w.Header().Set("content-type", "text/html")
			var t *template.Template
			if config.TemplateFile == "" {
				fmt.Println("Not implemented. please indicate TemplateIndexFile.")
				return
			}
			t, err = template.ParseFiles(config.Serve.TemplateIndexFile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.Execute(w, entries)
			if err != nil {
				log.Println(err)
			}
		} else {
			p := filepath.Join(config.ReportDir, req.URL.Path)
			b, err := ioutil.ReadFile(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			body := string(github_flavored_markdown.Markdown(b))
			var t *template.Template
			if config.Serve.TemplateBodyFile == "" {
				fmt.Println("Not implemented. please indicate TemplateBodyFile.")
				return
			}
			t, err = template.ParseFiles(config.Serve.TemplateBodyFile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.Execute(w, entry{
				Name: req.URL.Path,
				Body: template.HTML(body),
			})
		}
	})
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir(config.Serve.AssetsDir))))

	port := ":8080"
	url := "http://localhost" + port

	browser.OpenURL(url)
	http.ListenAndServe(port, nil)
}

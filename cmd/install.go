// Copyright © 2016 Wei-Ting Kuo <waitingkuo0527@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/waitingkuo/belt/utils"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		homeDir, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		rootPath := filepath.Join(homeDir, ".belt")

		// FIXME
		// move the store relative codes to store.go
		binPath := filepath.Join(rootPath, "bin")
		if err = os.MkdirAll(binPath, 0755); err != nil {
			panic(err)
		}

		if len(args) == 0 {
			os.Exit(0) // should define a better code
		}

		packageName := args[0]
		// FIXME
		// should build a better way to download packages
		// check how congo or brew work
		// missing verson control
		// move to another dir (pkg?)

		// consider GOARCH
		var rawurl string
		switch runtime.GOOS {
		case "linux":
			rawurl = "https://github.com/coreos/etcd/releases/download/v3.0.6/etcd-v3.0.6-linux-amd64.zip"
		case "darwin":
			rawurl = "https://github.com/coreos/etcd/releases/download/v3.0.6/etcd-v3.0.6-darwin-amd64.zip"
		default:
			fmt.Println(runtime.GOOS, "isn't supported")
		}
		if packageName == "etcd" {
			fmt.Println("Downloading ...")
			destPath, err := utils.Download(rawurl, "/tmp")
			if err != nil {
				panic(err)
			}
			fmt.Println("unziping", destPath)
			//fmt.Println("unzip ok")
			if err := utils.Unzip(destPath, "/tmp"); err != nil {
				panic(err)
			}
			dir := strings.TrimSuffix(destPath, filepath.Ext(destPath))
			os.Rename(filepath.Join(dir, "etcd"), filepath.Join(binPath, "etcd"))
			os.Rename(filepath.Join(dir, "etcdctl"), filepath.Join(binPath, "etcdctl"))

		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

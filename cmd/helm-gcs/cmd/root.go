// Copyright © 2018 Valentin Tjoncke <valtjo@gmail.com>
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
	"os"

	"cloud.google.com/go/storage"
	"github.com/lodotek/helm-gcs/pkg/gcs"
	"github.com/lodotek/helm-gcs/pkg/repo"
	"github.com/spf13/cobra"
)

var (
	gcsClient *storage.Client

	flagServiceAccount string
	flagDebug          bool
)

var rootCmd = &cobra.Command{
	Use:   "helm-gcs",
	Short: "Manage Helm repositories on Google Cloud Storage",
	Long:  ``,
}

// Execute executes the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		var err error
		gcsClient, err = gcs.NewClient(flagServiceAccount)
		if err != nil {
			panic(err)
		}
		if flagDebug {
			repo.Debug = true
		}
	})
	rootCmd.PersistentFlags().StringVar(&flagServiceAccount, "service-account", "", "service account to use for GCS")
	rootCmd.PersistentFlags().BoolVar(&flagDebug, "debug", false, "activate debug")
}

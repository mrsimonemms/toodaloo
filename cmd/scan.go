/*
Copyright © 2024 Simon Emms <simon@simonemms.com>

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

	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/config"
	"github.com/mrsimonemms/toodaloo/pkg/output"
	"github.com/mrsimonemms/toodaloo/pkg/scanner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var scanOpts struct {
	Glob        string
	GitFiles    bool
	IgnorePaths []string
	Output      string
	SavePath    string
	Tags        []string
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a project",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		logger.Log().Trace("Validating the input options")
		if _, ok := output.Types[scanOpts.Output]; !ok {
			return fmt.Errorf("unknown output type: %s", scanOpts.Output)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := &config.Config{
			IgnorePaths:      scanOpts.IgnorePaths,
			Output:           scanOpts.Output,
			SavePath:         scanOpts.SavePath,
			Tags:             scanOpts.Tags,
			WorkingDirectory: rootOpts.WorkingDirectory,
		}

		s, err := scanner.New(cfg)
		if err != nil {
			logger.Log().WithError(err).Error("Unable to load scanner")
			return err
		}

		files := args
		if scanOpts.GitFiles {
			if files, err = s.FindFilesByGit(); err != nil {
				return err
			}
		} else if len(args) == 0 {
			if files, err = s.FindFilesByGlob(scanOpts.Glob); err != nil {
				return err
			}
		}

		logger.Log().Info("Scanning for todos")
		result, err := s.ScanFiles(files)
		if err != nil {
			logger.Log().WithError(err).Error("Unable to scan files")
			return err
		}

		logger.Log().Info("Generating report")
		if err := output.Generate(scanOpts.Output, scanOpts.SavePath, result); err != nil {
			logger.Log().WithError(err).Error("Failed to generate report")
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	bindEnv("glob", "**/*")
	scanCmd.Flags().StringVar(&scanOpts.Glob, "glob", viper.GetString("glob"), "glob pattern - ignored if files provided as arguments")

	bindEnv("git-files", false)
	scanCmd.Flags().BoolVar(&scanOpts.GitFiles, "git-files", viper.GetBool("git-files"), "get files from the git tree")

	bindEnv("ignore-paths", []string{".git/**/*"})
	scanCmd.Flags().StringSliceVar(&scanOpts.IgnorePaths, "ignore-paths", viper.GetStringSlice("ignore-paths"), "ignore scanning these files")

	bindEnv("output", "yaml")
	scanCmd.Flags().StringVarP(&scanOpts.Output, "output", "o", viper.GetString("output"), "output type")

	bindEnv("save-path", ".toodaloo.yaml")
	scanCmd.Flags().StringVarP(&scanOpts.SavePath, "save-path", "s", viper.GetString("save-path"), `save report to path - use "-" to output to stdout`)

	bindEnv("tags", []string{"fixme", "todo", "@todo"})
	scanCmd.Flags().StringSliceVarP(&scanOpts.Tags, "tags", "t", viper.GetStringSlice("tags"), "todo tags")
}

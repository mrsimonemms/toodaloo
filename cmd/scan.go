/*
 * Copyright 2023 Simon Emms <simon@simonemms.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cmd

import (
	"fmt"

	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/config"
	"github.com/mrsimonemms/toodaloo/pkg/reports"
	"github.com/mrsimonemms/toodaloo/pkg/scanner"
	"github.com/spf13/cobra"
)

var scanOpts struct {
	Output     string
	ReportType string
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		report, ok := reports.Reports[scanOpts.ReportType]
		if !ok {
			logger.Log().WithField("report", scanOpts.ReportType).Fatal("Unknown report type")
			return fmt.Errorf("unknown report type: %s", scanOpts.ReportType)
		}

		cfg, err := config.NewConfigFromFile(rootCfg.WorkingDirectory, rootCfg.CfgFile)
		if err != nil {
			logger.Log().WithError(err).Error("Error reading config")
			return err
		}

		s, err := scanner.New(rootCfg.WorkingDirectory, cfg)
		if err != nil {
			logger.Log().WithError(err).Error("Unable to load scanner")
			return err
		}

		result, err := s.Exec()
		if err != nil {
			logger.Log().WithError(err).Error("Unable to scan files")
			return err
		}

		logger.Log().WithField("report-type", scanOpts.ReportType).Info("Generating report")
		return reports.Generate(scanOpts.Output, result, report)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&scanOpts.Output, "output", "o", ".toodaloo.yaml", "toodaloo report path")
	scanCmd.Flags().StringVarP(&scanOpts.ReportType, "report-type", "r", "toodaloo", "report type")
}

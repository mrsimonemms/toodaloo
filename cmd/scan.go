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
	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/config"
	"github.com/mrsimonemms/toodaloo/pkg/scanner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a project",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		return s.Exec()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

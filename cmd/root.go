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
	"os"
	"strings"

	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envvarPrefix = "TOODALOO"
)

var rootOpts struct {
	LogLevel         string
	WorkingDirectory string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toodaloo",
	Short: "Say goodbye to your todos",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := logger.SetLevel(rootOpts.LogLevel); err != nil {
			return err
		}
		logger.Logger.WithField("level", rootOpts.LogLevel).Trace("Setting log level")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func bindEnv(key string, defaultValue ...any) {
	envvarName := fmt.Sprintf("%s_%s", envvarPrefix, key)
	envvarName = strings.Replace(envvarName, "-", "_", -1)
	envvarName = strings.ToUpper(envvarName)

	err := viper.BindEnv(key, envvarName)
	cobra.CheckErr(err)

	for _, val := range defaultValue {
		viper.SetDefault(key, val)
	}
}

func init() {
	cwd, err := os.Getwd()
	cobra.CheckErr(err)

	bindEnv("directory", cwd)
	rootCmd.PersistentFlags().StringVarP(&rootOpts.WorkingDirectory, "directory", "d", viper.GetString("directory"), "working directory")

	bindEnv("log-level", logrus.InfoLevel)
	rootCmd.PersistentFlags().StringVarP(&rootOpts.LogLevel, "log-level", "l", viper.GetString("log-level"), fmt.Sprintf("log level: %s", logger.GetAllLevels()))
}

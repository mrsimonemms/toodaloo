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

package config

import (
	"os"
	"path/filepath"

	"github.com/mrsimonemms/golang-helpers/logger"
	"sigs.k8s.io/yaml"
)

func Init(directory, filename string) error {
	file := filepath.Join(directory, filename)
	l := logger.Log().WithField("filename", file)

	c, err := yaml.Marshal(defaultConfig)
	if err != nil {
		l.Error("Failed to generate default config")
		return err
	}

	l.Info("Saving config to file")

	if err := os.WriteFile(file, c, 0o644); err != nil {
		l.WithError(err).Error("Failed to save file")
		return err
	}

	return nil
}

func NewConfigFromFile(directory, filename string) (*Config, error) {
	file := filepath.Join(directory, filename)
	l := logger.Log().WithField("filename", file)

	l.Infof("Loading configuration from file")

	bytes, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if os.IsNotExist(err) {
		l.Info("Config file does not exist, using the default config")

		return &defaultConfig, nil
	}

	// Start from the default config
	config := defaultConfig
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

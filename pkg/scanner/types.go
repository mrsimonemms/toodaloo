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

package scanner

import (
	"os"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/config"
	"golang.org/x/exp/slices"
)

type Scan struct {
	config           *config.Config
	workingDirectory string
}

func (s *Scan) getListOfFiles() ([]string, error) {
	fileList := make([]string, 0)
	ignoreList := make([]string, 0)

	fsys := os.DirFS(s.workingDirectory)
	matches, err := doublestar.Glob(fsys, s.config.Glob)
	if err != nil {
		return nil, err
	}

	for _, i := range s.config.IgnorePaths {
		m, err := doublestar.Glob(fsys, i)
		if err != nil {
			return nil, err
		}

		ignoreList = append(ignoreList, m...)
	}

	for _, f := range matches {
		if !slices.Contains(ignoreList, f) {
			fileInfo, err := os.Stat(f)
			if err != nil {
				return nil, err
			}

			// Check path isn't a directory
			if !fileInfo.IsDir() {
				fileList = append(fileList, f)
			}
		}
	}

	return fileList, nil
}

func (s *Scan) Exec() error {
	files, err := s.getListOfFiles()
	if err != nil {
		logger.Log().WithError(err).Error("Failed to get list of files")
		return err
	}

	logger.Log().WithField("files", files).Debug("File list")
	logger.Log().WithField("file-count", len(files)).Info("Files found")

	return nil
}

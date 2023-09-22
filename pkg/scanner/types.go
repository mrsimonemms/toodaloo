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
	"fmt"
	"os"
	"path"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/config"
	"golang.org/x/sync/errgroup"
)

type Scan struct {
	config           *config.Config
	workingDirectory string
}

type ScanResult struct {
	File       string `json:"file"`
	LineNumber int    `json:"lineNumber"`
	Author     string `json:"author,omitempty"`
	Msg        string `json:"message,omitempty"`
}

func (s *Scan) getListOfFiles() ([]string, error) {
	fileList := make([]string, 0)
	// @todo(sje): add .toodaloorc.yaml in here
	ignoreList := make(map[string]bool, 0)

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

		for _, f := range m {
			ignoreList[f] = true
		}
	}

	for _, f := range matches {
		// Check file not in ignore list
		if _, inList := ignoreList[f]; !inList {
			filepath := path.Join(s.workingDirectory, f)
			// Now check it's not a directory
			fileInfo, err := os.Stat(filepath)
			if err != nil {
				return nil, err
			}

			// Check path isn't a directory
			if !fileInfo.IsDir() {
				fileList = append(fileList, filepath)
			}
		}
	}

	return fileList, nil
}

func (s *Scan) scanFilesForTodos(files []string) error {
	g := new(errgroup.Group)

	res := make([]ScanResult, 0)
	logger.Log().WithField("tags", s.config.Tags).Debug("Scanning files for todos")
	for _, file := range files {

		file := file
		l := logger.Log().WithField("file", file)

		g.Go(func() error {
			r, err := scanForTodos(l, file, s.config.Tags)
			if err != nil {
				return err
			}

			res = append(res, r...)

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	logger.Log().Debug("File scan finished")
	return nil
}

func (s *Scan) Exec() error {
	files, err := s.getListOfFiles()
	if err != nil {
		logger.Log().WithError(err).Error("Failed to get list of files")
		return err
	}

	logger.Log().WithField("files", files).Debug("File list")
	logger.Log().WithField("file-count", len(files)).Info("Files found")

	if err := s.scanFilesForTodos(files); err != nil {
		logger.Log().WithError(err).Error("Error scanning files")
		return err
	}
	return nil
}

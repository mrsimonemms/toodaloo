/*
 * Copyright 2024 Simon Emms <simon@simonemms.com>
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
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/config"
	"golang.org/x/sync/errgroup"
)

func New(cfg *config.Config) (*Scan, error) {
	return &Scan{
		config: cfg,
	}, nil
}

type Scan struct {
	config *config.Config
}

func (s *Scan) FindFilesByGit() ([]string, error) {
	files := make([]string, 0)

	repo, err := git.PlainOpen(s.config.WorkingDirectory)
	if err != nil {
		logger.Log().WithError(err).Error("Failed to open Git repo")
		return nil, err
	}

	h, err := repo.ResolveRevision("HEAD")
	if err != nil {
		logger.Log().WithError(err).Error("Failed to resolve Git revision")
		return nil, err
	}

	commit, err := repo.CommitObject(*h)
	if err != nil {
		logger.Log().WithError(err).Error("Failed to get the commit object")
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		logger.Log().WithError(err).Error("Failed to get commit tree")
		return nil, err
	}

	seen := make(map[plumbing.Hash]bool, 0)
	iter := object.NewTreeWalker(tree, true, seen)

	defer func() {
		logger.Log().Debug("Closing tree walker")
		iter.Close()
	}()

	// Based on file iterator, but without reading the file blob
	// to speed things up.
	// @link https://github.com/go-git/go-git/blob/302dddeda962e4bb3477a8e4080bc6f5a253e2bb/plumbing/object/file.go#L89,L107
	for {
		name, entry, err := iter.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Log().WithError(err).Error("Failed to iterate over files")
			return nil, err
		}

		if entry.Mode == filemode.Dir || entry.Mode == filemode.Submodule {
			continue
		}

		files = append(files, name)
	}

	return files, nil
}

func (s *Scan) FindFilesByGlob(glob string) ([]string, error) {
	fileList := make([]string, 0)
	ignoreList := make(map[string]bool, 0)

	fsys := os.DirFS(s.config.WorkingDirectory)
	matches, err := doublestar.Glob(fsys, glob)
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
			filepath := path.Join(s.config.WorkingDirectory, f)
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

func (s *Scan) scanFileForTodo(filename string) ([]Report, error) {
	l := logger.Log().WithField("file", filename)

	l.Debug("Scan starting")
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		l.Debug("Closing file")
		err = f.Close()
	}()

	res := make([]Report, 0)
	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++

		line := strings.TrimSpace(scanner.Text())

		l1 := l.WithField("line-number", lineNumber)
		l1.WithField("line", line)

		for _, tag := range s.config.Tags {
			regex := fmt.Sprintf(`^(\W+)(\s+)%s([\[\(](.*)[\]\)])?(:?(.*))?`, regexp.QuoteMeta(tag))

			l1.WithField("regex", regex).Debug("Searching in file")

			r, err := regexp.Compile(regex)
			if err != nil {
				return nil, err
			}

			matches := r.FindStringSubmatch(line)
			if len(matches) == 0 {
				continue
			}

			l1.WithField("matches", matches).Debug("Found matches")

			// @todo(sje): get the author from the Git history
			author := strings.TrimSpace(matches[4])
			msg := strings.TrimSpace(matches[6])

			// Remove working directory
			cleanFilename := strings.TrimPrefix(filename, s.config.WorkingDirectory)
			// Remove prefixed slash (may or may not be in working directory)
			cleanFilename = strings.TrimPrefix(cleanFilename, "/")

			res = append(res, Report{
				File:       cleanFilename,
				LineNumber: lineNumber,
				Author:     author,
				Msg:        msg,
			})
		}
	}

	l.WithField("total-lines", lineNumber).Debug("Scan ending")

	return res, err
}

func (s *Scan) ScanFiles(files []string) ([]Report, error) {
	g := new(errgroup.Group)

	res := make([]Report, 0)
	logger.Log().WithField("tags", s.config.Tags).Debug("Scanning files for todos")
	for _, file := range files {
		targetFile := file
		g.Go(func() error {
			r, err := s.scanFileForTodo(targetFile)
			if err != nil {
				return err
			}

			res = append(res, r...)

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	logger.Log().Debug("File scan finished")

	return res, nil
}

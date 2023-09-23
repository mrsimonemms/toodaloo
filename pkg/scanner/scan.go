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
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mrsimonemms/toodaloo/pkg/config"
	"github.com/sirupsen/logrus"
)

func New(workingDirectory string, cfg *config.Config) (*Scan, error) {
	return &Scan{
		config:           cfg,
		workingDirectory: workingDirectory,
	}, nil
}

func scanForTodos(l *logrus.Entry, filename string, tags []string) ([]ScanResult, error) {
	l.Debug("Scan starting")

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		l.Debug("Closing file")
		err = f.Close()
	}()

	res := make([]ScanResult, 0)

	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		for _, tag := range tags {
			regex := fmt.Sprintf("(^|\\s)%s(\\((.*)\\))?(:?(.*))?", regexp.QuoteMeta(tag))

			l = l.WithField("line", line)

			l.WithField("regex", regex).Debug("Searching in file")

			r, err := regexp.Compile(regex)
			if err != nil {
				return nil, err
			}

			matches := r.FindStringSubmatch(scanner.Text())
			l.WithField("matches", matches).Debug("Found matches")
			if len(matches) == 0 {
				continue
			}

			author := strings.TrimSpace(matches[3])
			msg := strings.TrimSpace(matches[5])

			res = append(res, ScanResult{
				File:       filename,
				LineNumber: line,
				Author:     author,
				Msg:        msg,
			})
		}

		line++
	}

	l.Debug("Scan ending")
	return res, err
}

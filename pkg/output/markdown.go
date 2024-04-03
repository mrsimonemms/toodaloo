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

package output

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrsimonemms/toodaloo/pkg/scanner"
)

type MarkdownOutput struct{}

// Convert to a markdown table
// @todo(sje): implement report
func (o MarkdownOutput) generate(report []scanner.Report) ([]byte, error) {
	res := [][]string{
		{"File", "Line Number", "Author", "Message"},
		{"---", "---", "---", "---"},
	}

	for _, item := range report {
		res = append(res, []string{
			item.File,
			strconv.Itoa(item.LineNumber),
			item.Author,
			item.Msg,
		})
	}

	output := strings.Join([]string{
		"# List of todos",
		"",
		fmt.Sprintf("> Generated by [Toodaloo](%s)", URL),
		"",
		"",
	}, "\n")
	for _, line := range res {
		line = append([]string{""}, line...)
		line = append(line, "")
		output += fmt.Sprintf("%s\n", strings.TrimSpace(strings.Join(line, " | ")))
	}

	return []byte(output), nil
}

func init() {
	// Register the report type
	Types["markdown"] = MarkdownOutput{}
}
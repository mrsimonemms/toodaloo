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
	"strconv"

	"github.com/mrsimonemms/toodaloo/pkg/scanner"
)

type MarkdownOutput struct{}

// Convert to a markdown table
// @todo(sje): implement report
func (o MarkdownOutput) Generate(report []scanner.Report) ([]byte, error) {
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

	return nil, nil
}

func init() {
	// Register the report type
	Types["markdown"] = MarkdownOutput{}
}

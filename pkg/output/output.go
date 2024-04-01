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
	"os"

	"github.com/mrsimonemms/toodaloo/pkg/scanner"
)

var Types map[string]Output = make(map[string]Output, 0)

type Output interface {
	generate([]scanner.Report) ([]byte, error)
}

func init() {
	// Register the output type
	Types["yaml"] = YamlOutput{}
}

func Generate(reportType, savePath string, report []scanner.Report) error {
	r, ok := Types[reportType]
	if !ok {
		return fmt.Errorf("unknown output type: %s", reportType)
	}

	output, err := r.generate(report)
	if err != nil {
		return err
	}

	if savePath == "-" {
		// Print to stdout
		fmt.Println(string(output))
		return nil
	}

	// Save to file
	return os.WriteFile(savePath, output, 0o644)
}

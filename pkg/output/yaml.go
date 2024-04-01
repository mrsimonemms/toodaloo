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
	"strings"

	"github.com/mrsimonemms/toodaloo/pkg/scanner"
	"sigs.k8s.io/yaml"
)

// Convert to YAML - this is the default report
type YamlOutput struct{}

// This is a basic output, just dumping the report to YAML
func (o YamlOutput) generate(res []scanner.Report) ([]byte, error) {
	y, err := yaml.Marshal(res)
	if err != nil {
		return nil, err
	}

	output := strings.Join([]string{
		"###",
		"# List of todos",
		fmt.Sprintf("# Generated by Toodaloo - %s", URL),
		"###",
		string(y),
	}, "\n")

	return []byte(output), nil
}

func init() {
	// Register the output type
	Types["yaml"] = YamlOutput{}
}

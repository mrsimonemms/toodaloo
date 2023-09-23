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

package reports

import (
	"github.com/mrsimonemms/toodaloo/pkg/scanner"
	"sigs.k8s.io/yaml"
)

type ToodalooReport struct{}

// This is a basic report, just dumping the report to YAML
func (t ToodalooReport) Generate(res []scanner.ScanResult) ([]byte, error) {
	return yaml.Marshal(res)
}

func init() {
	// Register the report type
	Reports["toodaloo"] = ToodalooReport{}
}

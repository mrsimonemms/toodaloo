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
	"os"

	"github.com/mrsimonemms/golang-helpers/logger"
	"github.com/mrsimonemms/toodaloo/pkg/scanner"
)

var Reports map[string]Report = make(map[string]Report, 0)

type Report interface {
	Generate([]scanner.ScanResult) ([]byte, error)
}

func Generate(filepath string, result []scanner.ScanResult, report Report) error {
	data, err := report.Generate(result)
	if err != nil {
		return err
	}

	logger.Log().WithField("file", filepath).Info("Writing report to file")
	return os.WriteFile(filepath, data, 0o644)
}
